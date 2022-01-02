package gui

import (
	"fmt"

	"github.com/jesseduffield/lazygit/pkg/gui/style"
)

func (gui *Gui) handleCreateResetMenu() error {
	red := style.FgRed

	nukeStr := "reset --hard HEAD && git clean -fd"
	if len(gui.State.Submodules) > 0 {
		nukeStr = fmt.Sprintf("%s (%s)", nukeStr, gui.Tr.LcAndResetSubmodules)
	}

	menuItems := []*menuItem{
		{
			displayStrings: []string{
				gui.Tr.LcDiscardAllChangesToAllFiles,
				red.Sprint(nukeStr),
			},
			onPress: func() error {
				if err := gui.GitCommand.WithSpan(gui.Tr.Spans.NukeWorkingTree).WorkingTree.ResetAndClean(); err != nil {
					return gui.surfaceError(err)
				}

				return gui.refreshSidePanels(refreshOptions{mode: ASYNC, scope: []RefreshableView{FILES}})
			},
		},
		{
			displayStrings: []string{
				gui.Tr.LcDiscardAnyUnstagedChanges,
				red.Sprint("git checkout -- ."),
			},
			onPress: func() error {
				if err := gui.GitCommand.WithSpan(gui.Tr.Spans.DiscardUnstagedFileChanges).WorkingTree.DiscardAnyUnstagedFileChanges(); err != nil {
					return gui.surfaceError(err)
				}

				return gui.refreshSidePanels(refreshOptions{mode: ASYNC, scope: []RefreshableView{FILES}})
			},
		},
		{
			displayStrings: []string{
				gui.Tr.LcDiscardUntrackedFiles,
				red.Sprint("git clean -fd"),
			},
			onPress: func() error {
				if err := gui.GitCommand.WithSpan(gui.Tr.Spans.RemoveUntrackedFiles).WorkingTree.RemoveUntrackedFiles(); err != nil {
					return gui.surfaceError(err)
				}

				return gui.refreshSidePanels(refreshOptions{mode: ASYNC, scope: []RefreshableView{FILES}})
			},
		},
		{
			displayStrings: []string{
				gui.Tr.LcSoftReset,
				red.Sprint("git reset --soft HEAD"),
			},
			onPress: func() error {
				if err := gui.GitCommand.WithSpan(gui.Tr.Spans.SoftReset).WorkingTree.ResetSoft("HEAD"); err != nil {
					return gui.surfaceError(err)
				}

				return gui.refreshSidePanels(refreshOptions{mode: ASYNC, scope: []RefreshableView{FILES}})
			},
		},
		{
			displayStrings: []string{
				"mixed reset",
				red.Sprint("git reset --mixed HEAD"),
			},
			onPress: func() error {
				if err := gui.GitCommand.WithSpan(gui.Tr.Spans.MixedReset).WorkingTree.ResetMixed("HEAD"); err != nil {
					return gui.surfaceError(err)
				}

				return gui.refreshSidePanels(refreshOptions{mode: ASYNC, scope: []RefreshableView{FILES}})
			},
		},
		{
			displayStrings: []string{
				gui.Tr.LcHardReset,
				red.Sprint("git reset --hard HEAD"),
			},
			onPress: func() error {
				if err := gui.GitCommand.WithSpan(gui.Tr.Spans.HardReset).WorkingTree.ResetHard("HEAD"); err != nil {
					return gui.surfaceError(err)
				}

				return gui.refreshSidePanels(refreshOptions{mode: ASYNC, scope: []RefreshableView{FILES}})
			},
		},
	}

	return gui.createMenu("", menuItems, createMenuOptions{showCancel: true})
}
