package page

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/slok/brigadeterm/pkg/controller"
)

const (
	// ProjectListPageName is the name that identifies thi projectList page.
	ProjectListPageName = "projectlist"
)

// ProjectList is the main page where the project list will be available.
type ProjectList struct {
	controller controller.Controller
	router     *Router

	// page layout
	layout tview.Primitive

	// components
	projectsTable *tview.Table

	registerPageOnce sync.Once
}

// NewProjectList returns a new project list.
func NewProjectList(controller controller.Controller, router *Router) *ProjectList {
	p := &ProjectList{
		controller: controller,
		router:     router,
	}
	p.createComponents()
	return p
}

// Register satisfies Page interface.
func (p *ProjectList) Register(pages *tview.Pages) {
	p.registerPageOnce.Do(func() {
		pages.AddPage(ProjectListPageName, p.layout, true, false)
	})
}

// BeforeLoad satisfies Page interface.
func (p *ProjectList) BeforeLoad() {
}

// Refresh will refresh all the page data.
func (p *ProjectList) Refresh() {
	ctx := p.controller.ProjectListPageContext()
	// TODO: check error.
	p.fill(ctx)
}

// createComponents will create all the layout components.
func (p *ProjectList) createComponents() {
	// Set up columns
	p.projectsTable = tview.NewTable().
		SetSelectable(true, false)
	p.projectsTable.
		SetBorder(true).
		SetTitle("Projects")

	// Create the layout.
	p.layout = tview.NewFlex().
		AddItem(p.projectsTable, 0, 1, true)
}

// FillBrigadeProjectList
func (p *ProjectList) fill(ctx *controller.ProjectListPageContext) {
	// Clear other widgets.
	p.projectsTable.Clear()

	// Set header.
	p.projectsTable.SetCell(0, 0, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 1, &tview.TableCell{Text: "Last build type", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 2, &tview.TableCell{Text: "Last build version", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 3, &tview.TableCell{Text: "Last build time", Align: tview.AlignCenter, Color: tcell.ColorYellow})

	projectNameIDIndex := map[string]string{}

	// Set body.
	rowPosition := 1
	for _, project := range ctx.Projects {

		// Select row color.
		color := tcell.ColorWhite
		if !project.LastBuild.Running {
			color = tcell.ColorRed
			if project.LastBuild.FinishedOK {
				color = tcell.ColorGreen
			}
		}

		// Set the index so we can get the project ID on selection.
		projectNameIDIndex[project.Name] = project.ID

		p.projectsTable.SetCell(rowPosition, 0, &tview.TableCell{Text: project.Name, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 1, &tview.TableCell{Text: project.LastBuild.EventType, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 2, &tview.TableCell{Text: project.LastBuild.Version, Align: tview.AlignLeft, Color: color})
		timeAgo := time.Now().Sub(project.LastBuild.Started)
		p.projectsTable.SetCell(rowPosition, 3, &tview.TableCell{Text: fmt.Sprintf("%v ago", timeAgo), Align: tview.AlignLeft, Color: color})

		rowPosition++
	}

	// Set selectable to call our jobs.
	p.projectsTable.SetSelectedFunc(func(row, column int) {
		// Get project ID cell and from commit the build ID.
		cell := p.projectsTable.GetCell(row, 0)
		projectID := projectNameIDIndex[cell.Text]
		// Load build list page.
		p.router.LoadProjectBuildList(projectID)
	})
}
