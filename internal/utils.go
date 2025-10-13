package internal

import (
	"fmt"
	"strconv"
	"strings"

	core "github.com/statloc/core"
)

func GetTable(items map[string]*core.TableItem) string {
    // im sorry for that
    maxTitleLength, maxLOCLength, maxFilesLength := 5, 3, 5
    for title, item := range items {
        if len(title) > maxTitleLength {
            maxTitleLength = len(title)
        }

        LOC := strconv.FormatUint(item.LOC, 10)
        if len(LOC) > maxLOCLength {
            maxLOCLength = len(LOC)
        }

        files := strconv.FormatUint(item.Files, 10)
        if len(files) > maxFilesLength {
            maxFilesLength = len(files)
        }
    }

    separator := fmt.Sprintf(
        "+-%s-+-%s-+-%s-+\n",
        strings.Repeat("-", maxTitleLength),
        strings.Repeat("-", maxLOCLength),
        strings.Repeat("-", maxFilesLength),
    )

    result := fmt.Sprint(
        separator,
        fmt.Sprintf(
            "| Title%s | LOC%s | Files%s |\n",
            strings.Repeat(" ", maxTitleLength - 5),
            strings.Repeat(" ", maxLOCLength - 3),
            strings.Repeat(" ", maxFilesLength - 5),
        ),
        separator,
    )

    for title, item := range items {
        result += fmt.Sprintf(
            "| %s%s | %d%s | %d%s |\n",
            title, strings.Repeat(" ", maxTitleLength - len(title)),
            item.LOC, strings.Repeat(" ", maxLOCLength - len(strconv.FormatUint(item.LOC, 10))),
            item.Files, strings.Repeat(" ", maxFilesLength - len(strconv.FormatUint(item.Files, 10))),
        )
    }

    result += separator
    return result
}
