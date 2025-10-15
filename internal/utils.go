package internal

import (
	"fmt"
	"strconv"
	"strings"

	core "github.com/statloc/core"
)

func GetTable(
    items       map[string]*core.TableItem,
    titleLength int,
    LOCLength   int,
    filesLength int,
) string {
    // im sorry for that
    column1width, column2width, column3width := titleLength, LOCLength, filesLength
    for title, item := range items {
        if len(title) > int(column1width) {
            column1width = len(title)
        }

        LOC := strconv.FormatUint(item.LOC, 10)
        if len(LOC) > column2width {
            column2width= len(LOC)
        }

        files := strconv.FormatUint(item.Files, 10)
        if len(files) > column3width {
            column3width= len(files)
        }
    }

    separator := fmt.Sprintf(
        "+-%s-+-%s-+-%s-+\n",
        strings.Repeat("-", column1width),
        strings.Repeat("-", column2width),
        strings.Repeat("-", column3width),
    )

    result := fmt.Sprint(
        separator,
        fmt.Sprintf(
            "| Title%s | LOC%s | Files%s |\n",
            strings.Repeat(" ", column1width - titleLength),
            strings.Repeat(" ", column2width - LOCLength),
            strings.Repeat(" ", column3width - filesLength),
        ),
        separator,
    )

    for title, item := range items {
        result += fmt.Sprintf(
            "| %s%s | %d%s | %d%s |\n",
            title, strings.Repeat(" ", column1width - len(title)),
            item.LOC, strings.Repeat(" ", column2width - len(strconv.FormatUint(item.LOC, 10))),
            item.Files, strings.Repeat(" ", column3width - len(strconv.FormatUint(item.Files, 10))),
        )
    }

    result += separator
    return result
}
