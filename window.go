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

func NewWindow(tp WindowType) *Window {
        ret := C.gtk_window_new(C.GtkWindowType(tp))
        if ret == nil {
                return nil
        }
        return ToWindow(unsafe.Pointer(ret))
}

func (win *Window) GetTitle() string {
        ret := C.gtk_window_get_title(win.c)
        if ret == nil {
                return ""
        }
        s := C.GoString((*C.char)(ret))

        return s
}

func (w *Window) SetWmclass(wmclass_name, wmclass_class string) {
        s0 := _GString(wmclass_name)
        s1 := _GString(wmclass_class)
        defer _GFree(unsafe.Pointer(s0))
        defer _GFree(unsafe.Pointer(s1))

        C.gtk_window_set_wmclass(w.c, s0, s1)
}

func (w *Window) SetRole(role string) {
        s0 := _GString(role)
        defer _GFree(unsafe.Pointer(s0))
        C.gtk_window_set_role(w.c, s0)
}

func (w *Window) GetRole() string {
        s0 := C.gtk_window_get_role(w.c)
        if s0 == nil {
                return ""
        }
        s := C.GoString((*C.char)(s0))

        return s
}

func (w *Window) SetStartupId(role string) {
        s0 := _GString(role)
        defer _GFree(unsafe.Pointer(s0))
        C.gtk_window_set_startup_id(w.c, s0)
}

func (win *Window) SetTitle(t string) {
        title := _GString(t)
        defer _GFree(unsafe.Pointer(title))
        C.gtk_window_set_title(win.c, title)
}

// void       gtk_window_add_accel_group          (GtkWindow           *window,
//                         GtkAccelGroup       *accel_group);
// void       gtk_window_remove_accel_group       (GtkWindow           *window,
//                         GtkAccelGroup       *accel_group);
// void       gtk_window_set_position             (GtkWindow           *window,
//                         GtkWindowPosition    position);
// gboolean   gtk_window_activate_focus           (GtkWindow           *window);
// void       gtk_window_set_focus                (GtkWindow           *window,
//                         GtkWidget           *focus);
// GtkWidget *gtk_window_get_focus                (GtkWindow           *window);
// void       gtk_window_set_default              (GtkWindow           *window,
//                         GtkWidget           *default_widget);
// GtkWidget *gtk_window_get_default_widget       (GtkWindow           *window);
// gboolean   gtk_window_activate_default         (GtkWindow           *window);

// void       gtk_window_set_transient_for        (GtkWindow           *window,
//                         GtkWindow           *parent);
// GtkWindow *gtk_window_get_transient_for        (GtkWindow           *window);
// GDK_AVAILABLE_IN_3_4
// void       gtk_window_set_attached_to          (GtkWindow           *window,
//                                                 GtkWidget           *attach_widget);
// GDK_AVAILABLE_IN_3_4
// GtkWidget *gtk_window_get_attached_to          (GtkWindow           *window);
// GDK_DEPRECATED_IN_3_8_FOR(gtk_widget_set_opacity)
// void       gtk_window_set_opacity              (GtkWindow           *window,
//                         gdouble              opacity);
// GDK_DEPRECATED_IN_3_8_FOR(gtk_widget_get_opacity)
// gdouble    gtk_window_get_opacity              (GtkWindow           *window);
// void       gtk_window_set_type_hint            (GtkWindow           *window,
//                         GdkWindowTypeHint    hint);
// GdkWindowTypeHint gtk_window_get_type_hint     (GtkWindow           *window);
// void       gtk_window_set_skip_taskbar_hint    (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_skip_taskbar_hint    (GtkWindow           *window);
// void       gtk_window_set_skip_pager_hint      (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_skip_pager_hint      (GtkWindow           *window);
// void       gtk_window_set_urgency_hint         (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_urgency_hint         (GtkWindow           *window);
// void       gtk_window_set_accept_focus         (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_accept_focus         (GtkWindow           *window);
// void       gtk_window_set_focus_on_map         (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_focus_on_map         (GtkWindow           *window);
// void       gtk_window_set_destroy_with_parent  (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_destroy_with_parent  (GtkWindow           *window);
// GDK_AVAILABLE_IN_3_4
// void       gtk_window_set_hide_titlebar_when_maximized (GtkWindow   *window,
//                                                         gboolean     setting);
// GDK_AVAILABLE_IN_3_4
// gboolean   gtk_window_get_hide_titlebar_when_maximized (GtkWindow   *window);
// void       gtk_window_set_mnemonics_visible    (GtkWindow           *window,
//                                                 gboolean             setting);
// gboolean   gtk_window_get_mnemonics_visible    (GtkWindow           *window);
// GDK_AVAILABLE_IN_3_2
// void       gtk_window_set_focus_visible        (GtkWindow           *window,
//                                                 gboolean             setting);
// GDK_AVAILABLE_IN_3_2
// gboolean   gtk_window_get_focus_visible        (GtkWindow           *window);

// void       gtk_window_set_resizable            (GtkWindow           *window,
//                                                 gboolean             resizable);
// gboolean   gtk_window_get_resizable            (GtkWindow           *window);

// void       gtk_window_set_gravity              (GtkWindow           *window,
//                                                 GdkGravity           gravity);
// GdkGravity gtk_window_get_gravity              (GtkWindow           *window);

// void       gtk_window_set_geometry_hints       (GtkWindow           *window,
//                         GtkWidget           *geometry_widget,
//                         GdkGeometry         *geometry,
//                         GdkWindowHints       geom_mask);

// void       gtk_window_set_screen           (GtkWindow       *window,
//                         GdkScreen       *screen);
// GdkScreen* gtk_window_get_screen           (GtkWindow       *window);

// gboolean   gtk_window_is_active                (GtkWindow           *window);
// gboolean   gtk_window_has_toplevel_focus       (GtkWindow           *window);

// void       gtk_window_set_decorated            (GtkWindow *window,
//                                                 gboolean   setting);
// gboolean   gtk_window_get_decorated            (GtkWindow *window);
// void       gtk_window_set_deletable            (GtkWindow *window,
//                                                 gboolean   setting);
// gboolean   gtk_window_get_deletable            (GtkWindow *window);

// void       gtk_window_set_icon_list                (GtkWindow  *window,
//                                                     GList      *list);
// GList*     gtk_window_get_icon_list                (GtkWindow  *window);
// void       gtk_window_set_icon                     (GtkWindow  *window,
//                                                     GdkPixbuf  *icon);
// void       gtk_window_set_icon_name                (GtkWindow   *window,
//                             const gchar *name);
// gboolean   gtk_window_set_icon_from_file           (GtkWindow   *window,
//                             const gchar *filename,
//                             GError     **err);
// GdkPixbuf* gtk_window_get_icon                     (GtkWindow  *window);
// const gchar * gtk_window_get_icon_name             (GtkWindow  *window);
// void       gtk_window_set_default_icon_list        (GList      *list);
// GList*     gtk_window_get_default_icon_list        (void);
// void       gtk_window_set_default_icon             (GdkPixbuf  *icon);
// void       gtk_window_set_default_icon_name        (const gchar *name);
// const gchar * gtk_window_get_default_icon_name     (void);
// gboolean   gtk_window_set_default_icon_from_file   (const gchar *filename,
//                             GError     **err);

// void       gtk_window_set_auto_startup_notification (gboolean setting);
/* If window is set modal, input will be grabbed when show and released when hide */
// void       gtk_window_set_modal      (GtkWindow *window,
//                       gboolean   modal);
// gboolean   gtk_window_get_modal      (GtkWindow *window);
// GList*     gtk_window_list_toplevels (void);
// void       gtk_window_set_has_user_ref_count (GtkWindow *window,
//                                               gboolean   setting);

// void     gtk_window_add_mnemonic          (GtkWindow       *window,
//                        guint            keyval,
//                        GtkWidget       *target);
// void     gtk_window_remove_mnemonic       (GtkWindow       *window,
//                        guint            keyval,
//                        GtkWidget       *target);
// gboolean gtk_window_mnemonic_activate     (GtkWindow       *window,
//                        guint            keyval,
//                        GdkModifierType  modifier);
// void     gtk_window_set_mnemonic_modifier (GtkWindow       *window,
//                        GdkModifierType  modifier);
// GdkModifierType gtk_window_get_mnemonic_modifier (GtkWindow *window);

// gboolean gtk_window_activate_key          (GtkWindow        *window,
//                        GdkEventKey      *event);
// gboolean gtk_window_propagate_key_event   (GtkWindow        *window,
//                        GdkEventKey      *event);

// void     gtk_window_present            (GtkWindow *window);
// void     gtk_window_present_with_time  (GtkWindow *window,
//                         guint32    timestamp);
// void     gtk_window_iconify       (GtkWindow *window);
// void     gtk_window_deiconify     (GtkWindow *window);
// void     gtk_window_stick         (GtkWindow *window);
// void     gtk_window_unstick       (GtkWindow *window);
// void     gtk_window_maximize      (GtkWindow *window);
// void     gtk_window_unmaximize    (GtkWindow *window);
// void     gtk_window_fullscreen    (GtkWindow *window);
// void     gtk_window_unfullscreen  (GtkWindow *window);
// void     gtk_window_set_keep_above    (GtkWindow *window, gboolean setting);
// void     gtk_window_set_keep_below    (GtkWindow *window, gboolean setting);

// void gtk_window_begin_resize_drag (GtkWindow     *window,
//                                    GdkWindowEdge  edge,
//                                    gint           button,
//                                    gint           root_x,
//                                    gint           root_y,
//                                    guint32        timestamp);
// void gtk_window_begin_move_drag   (GtkWindow     *window,
//                                    gint           button,
//                                    gint           root_x,
//                                    gint           root_y,
//                                    guint32        timestamp);
/* Set initial default size of the window (does not constrain user
 * resize operations)
 */
// void     gtk_window_set_default_size (GtkWindow   *window,
//                                       gint         width,
//                                       gint         height);
func (w *Window) SetDefaultSize(width, height int32) {
        C.gtk_window_set_default_size(w.c, C.gint(width), C.gint(height))
}

// void     gtk_window_get_default_size (GtkWindow   *window,
//                                       gint        *width,
//                                       gint        *height);
// void     gtk_window_resize           (GtkWindow   *window,
//                                       gint         width,
//                                       gint         height);
// void     gtk_window_get_size         (GtkWindow   *window,
//                                       gint        *width,
//                                       gint        *height);
// void     gtk_window_move             (GtkWindow   *window,
//                                       gint         x,
//                                       gint         y);
// void     gtk_window_get_position     (GtkWindow   *window,
//                                       gint        *root_x,
//                                       gint        *root_y);
// gboolean gtk_window_parse_geometry   (GtkWindow   *window,
//                                       const gchar *geometry);

// void gtk_window_set_default_geometry (GtkWindow *window,
//                                       gint       width,
//                                       gint       height);
// void gtk_window_resize_to_geometry   (GtkWindow *window,
//                                       gint       width,
//                                       gint       height);

// GtkWindowGroup *gtk_window_get_group (GtkWindow   *window);
// gboolean gtk_window_has_group        (GtkWindow   *window);
/* Ignore this unless you are writing a GUI builder */
// void     gtk_window_reshow_with_initial_size (GtkWindow *window);

// GtkWindowType gtk_window_get_window_type     (GtkWindow     *window);
// GtkApplication *gtk_window_get_application      (GtkWindow          *window);
// void            gtk_window_set_application      (GtkWindow          *window,
//                                                  GtkApplication     *application);
// void     gtk_window_set_has_resize_grip    (GtkWindow    *window,
//                                             gboolean      value);
// gboolean gtk_window_get_has_resize_grip    (GtkWindow    *window);
// gboolean gtk_window_resize_grip_is_visible (GtkWindow    *window);
// gboolean gtk_window_get_resize_grip_area   (GtkWindow    *window,
//                                             GdkRectangle *rect);
func (b *Window) CType() *C.GtkWindow {
        return b.c
}
func (w *Window) Show() {
        C.gtk_widget_show(ToWidget(unsafe.Pointer(w.c)).c)
}

func (w *Window) ShowAll() {
        C.gtk_widget_show_all(ToWidget(unsafe.Pointer(w.c)).c)
}

func (w *Window) Connect(sig string, f interface{}) {
        csignal := _GString(sig)
        goclosure := C.gGoclosureNew(unsafe.Pointer(&f), nil)
        C.g_signal_connect_closure(C.gpointer(unsafe.Pointer(w.c)),
                csignal, (*C.GClosure)(unsafe.Pointer(goclosure)), 0)
        _GFree(unsafe.Pointer(csignal))
}

func NewWindowGroup() *WindowGroup {
        ret := C.gtk_window_group_new()
        if ret == nil {
                return nil
        }
        return &WindowGroup{ret}
}

func (w *WindowGroup) AddWindow(win *Window) {
        C.gtk_window_group_add_window(w.c, win.c)
}

// void             gtk_window_group_remove_window (GtkWindowGroup     *window_group,
//                              GtkWindow          *window);
// GList *          gtk_window_group_list_windows  (GtkWindowGroup     *window_group);

// GtkWidget *      gtk_window_group_get_current_grab (GtkWindowGroup *window_group);
// GtkWidget *      gtk_window_group_get_current_device_grab (GtkWindowGroup *window_group,
//                                                       GdkDevice      *device);
