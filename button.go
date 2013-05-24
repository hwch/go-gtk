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

func NewButton() *Button {
        bt := C.gtk_button_new()
        if bt == nil {
                return nil
        }
        return ToButton(unsafe.Pointer(bt))
}

func NewButtonWithLabel(s string) *Button {
        l := _GString(s)
        defer _GFree(unsafe.Pointer(l))
        bt := C.gtk_button_new_with_label(l)
        if bt == nil {
                return nil
        }
        return ToButton(unsafe.Pointer(bt))
}

func (b *Button) SetLabel(label string) {
        l := _GString(label)
        defer _GFree(unsafe.Pointer(l))
        C.gtk_button_set_label(b.c, l)
}

func (b *Button) CType() *C.GtkButton {
        return b.c
}

func (b *Button) Connect(sig string, f interface{}) {
        csignal := _GString(sig)
        goclosure := C.gGoclosureNew(unsafe.Pointer(&f), nil)
        C.g_signal_connect_closure(C.gpointer(unsafe.Pointer(b.c)),
                csignal, (*C.GClosure)(unsafe.Pointer(goclosure)), 0)
        _GFree(unsafe.Pointer(csignal))
}

// GtkWidget*     gtk_button_new_from_stock    (const gchar    *stock_id);
func NewButtonFromStock(stock_id string) *Button {
        s := _GString(stock_id)
        defer _GFree(unsafe.Pointer(s))

        ret := C.gtk_button_new_from_stock(s)
        if ret != nil {
                return ToButton(unsafe.Pointer(ret))
        }

        return nil
}

// GtkWidget*     gtk_button_new_with_mnemonic (const gchar    *label);
func NewButtonWithMnemonic(label string) *Button {
        s := _GString(label)
        defer _GFree(unsafe.Pointer(s))

        ret := C.gtk_button_new_with_mnemonic(s)
        if ret != nil {
                return ToButton(unsafe.Pointer(ret))
        }

        return nil
}

// void           gtk_button_clicked           (GtkButton      *button);
func (b *Button) Clicked() {
        C.gtk_button_clicked(b.c)
}

// GDK_DEPRECATED
// void           gtk_button_pressed           (GtkButton      *button);
// GDK_DEPRECATED
// void           gtk_button_released          (GtkButton      *button);
// GDK_DEPRECATED
// void           gtk_button_enter             (GtkButton      *button);
// GDK_DEPRECATED
// void           gtk_button_leave             (GtkButton      *button);

// void                  gtk_button_set_relief         (GtkButton      *button,
// 						     GtkReliefStyle  newstyle);
func (b *Button) SetRelief(newstyle ReliefStyle) {
        C.gtk_button_set_relief(b.c, C.GtkReliefStyle(newstyle))
}

// GtkReliefStyle        gtk_button_get_relief         (GtkButton      *button);
func (b *Button) GetRelief() ReliefStyle {
        return ReliefStyle(C.gtk_button_get_relief(b.c))
}

// const gchar *         gtk_button_get_label          (GtkButton      *button);

// void                  gtk_button_set_use_underline  (GtkButton      *button,
// 						     gboolean        use_underline);
// gboolean              gtk_button_get_use_underline  (GtkButton      *button);
// void                  gtk_button_set_use_stock      (GtkButton      *button,
// 						     gboolean        use_stock);
// gboolean              gtk_button_get_use_stock      (GtkButton      *button);
// void                  gtk_button_set_focus_on_click (GtkButton      *button,
// 						     gboolean        focus_on_click);
// gboolean              gtk_button_get_focus_on_click (GtkButton      *button);
// void                  gtk_button_set_alignment      (GtkButton      *button,
// 						     gfloat          xalign,
// 						     gfloat          yalign);
// void                  gtk_button_get_alignment      (GtkButton      *button,
// 						     gfloat         *xalign,
// 						     gfloat         *yalign);
// void                  gtk_button_set_image          (GtkButton      *button,
// 					             GtkWidget      *image);
// GtkWidget*            gtk_button_get_image          (GtkButton      *button);
// void                  gtk_button_set_image_position (GtkButton      *button,
// 						     GtkPositionType position);
// GtkPositionType       gtk_button_get_image_position (GtkButton      *button);
// GDK_AVAILABLE_IN_3_6
// void                  gtk_button_set_always_show_image (GtkButton   *button,
//                                                         gboolean     always_show);
// GDK_AVAILABLE_IN_3_6
// gboolean              gtk_button_get_always_show_image (GtkButton   *button);

// GdkWindow*            gtk_button_get_event_window   (GtkButton      *button);
