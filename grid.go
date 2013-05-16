package gtk

/*
#cgo pkg-config: gobject-2.0 gtk+-3.0
#include <glib-object.h>
#include <gtk/gtk.h>
#include "gobj.h"
*/
import "C"

import (
        "unsafe"
)

func NewGrid() *Grid {
        ret := C.gtk_grid_new()
        if ret == nil {
                return nil
        }
        return ToGrid(unsafe.Pointer(ret))
}

func (g *Grid) Attach(child *Widget,
        left, top, width, height int32) {
        C.gtk_grid_attach(g.c, child.c, C.gint(left),
                C.gint(top), C.gint(width), C.gint(height))
}

func (g *Grid) CType() *C.GtkGrid {
        return g.c
}

// void       gtk_grid_attach_next_to         (GtkGrid         *grid,
//                                             GtkWidget       *child,
//                                             GtkWidget       *sibling,
//                                             GtkPositionType  side,
//                                             gint             width,
//                                             gint             height);
func (g *Grid) AttachNextTo(child, sibling *Widget,
        side PositionType, width, height int32) {
        C.gtk_grid_attach_next_to(g.c, child.c, sibling.c,
                C.GtkPositionType(side), C.gint(width), C.gint(height))
}

// GDK_AVAILABLE_IN_3_2
// GtkWidget *gtk_grid_get_child_at           (GtkGrid         *grid,
//                                             gint             left,
//                                             gint             top);
func (g *Grid) GetChildAt(left, top int32) *Widget {
        ret := C.gtk_grid_get_child_at(g.c, C.gint(left), C.gint(top))
        if ret == nil {
                return nil
        }
        return &Widget{ret}
}

// GDK_AVAILABLE_IN_3_2
// void       gtk_grid_insert_row             (GtkGrid         *grid,
//                                             gint             position);
func (g *Grid) InsertRow(position int32) {
        C.gtk_grid_insert_row(g.c, C.gint(position))
}

// GDK_AVAILABLE_IN_3_2
// void       gtk_grid_insert_column          (GtkGrid         *grid,
//                                             gint             position);
func (g *Grid) InsertColumn(position int32) {
        C.gtk_grid_insert_column(g.c, C.gint(position))
}

// GDK_AVAILABLE_IN_3_2
// void       gtk_grid_insert_next_to         (GtkGrid         *grid,
//                                             GtkWidget       *sibling,
//                                             GtkPositionType  side);
func (g *Grid) InsertNextTo(sibling *Widget, side PositionType) {
        C.gtk_grid_insert_next_to(g.c, sibling.c, C.GtkPositionType(side))
}

// void       gtk_grid_set_row_homogeneous    (GtkGrid         *grid,
//                                             gboolean         homogeneous);
func (g *Grid) SetRowHomogeneous(homogeneous bool) {
        var b C.gboolean = 0
        if homogeneous {
                b = 1
        }
        C.gtk_grid_set_row_homogeneous(g.c, b)
}

// gboolean   gtk_grid_get_row_homogeneous    (GtkGrid         *grid);
func (g *Grid) GetRowHomogeneous() bool {
        return C.gtk_grid_get_row_homogeneous(g.c) != 0
}

// void       gtk_grid_set_row_spacing        (GtkGrid         *grid,
//                                             guint            spacing);
func (g *Grid) SetRowSpacing(spacing uint32) {
        C.gtk_grid_set_row_spacing(g.c, C.guint(spacing))
}

// guint      gtk_grid_get_row_spacing        (GtkGrid         *grid);
func (g *Grid) GetRowSpacing() uint32 {
        return uint32(C.gtk_grid_get_row_spacing(g.c))
}

// void       gtk_grid_set_column_homogeneous (GtkGrid         *grid,
//                                             gboolean         homogeneous);
func (g *Grid) SetColumnHomogeneous(homogeneous bool) {
        var b C.gboolean = 0
        if homogeneous {
                b = 1
        }
        C.gtk_grid_set_column_homogeneous(g.c, b)
}

// gboolean   gtk_grid_get_column_homogeneous (GtkGrid         *grid);
func (g *Grid) GetColumnHomogeneous() bool {
        return C.gtk_grid_get_column_homogeneous(g.c) != 0
}

// void       gtk_grid_set_column_spacing     (GtkGrid         *grid,
//                                             guint            spacing);
func (g *Grid) SetColumnSpacing(spacing uint32) {
        C.gtk_grid_set_column_spacing(g.c, C.guint(spacing))
}

// guint      gtk_grid_get_column_spacing     (GtkGrid         *grid);
func (g *Grid) GetColumnSpacing() uint32 {
        return uint32(C.gtk_grid_get_column_spacing(g.c))
}
