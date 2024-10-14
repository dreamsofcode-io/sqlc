package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func newWriter() *tabwriter.Writer {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tClass\tLevel\tCreated At\tUpdated At")
	return w
}

//
// func printPlayer(x player.Player) {
// 	w := newWriter()
// 	defer w.Flush()
//
// 	printPlayerLine(w, x)
// }
//
// func printPlayerLine(w io.Writer, player player.Player) {
// 	fmt.Fprintf(w,
// 		"%d\t%s\t%s\t%d\t%s\t%s\n",
// 		player.ID,
// 		player.Name,
// 		player.Class,
// 		player.Level,
// 		player.CreatedAt.UTC().Format(time.RFC3339),
// 		player.UpdatedAt.UTC().Format(time.RFC3339),
// 	)
// }
//
// func printPlayers(xs []player.Player) {
// 	w := newWriter()
// 	defer w.Flush()
//
// 	for _, x := range xs {
// 		printPlayerLine(w, x)
// 	}
// }
