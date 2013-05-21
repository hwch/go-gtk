package gtk

/*
#cgo pkg-config: gobject-2.0 gtk+-3.0
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <glib-object.h>
#include <gtk/gtk.h>
#include <gtk/gtk-a11y.h>
#include "gobj.h"

GType _g_param_spec_type(GParamSpec *pspec) {
        return G_PARAM_SPEC_TYPE(pspec);
}

GType _g_param_spec_value_type(GParamSpec *pspec) {
        return G_PARAM_SPEC_VALUE_TYPE(pspec);
}

GType _g_object_type(GObject *obj) {
        return G_OBJECT_TYPE(obj);
}

GType _g_value_type(GValue *val) {
        return G_VALUE_TYPE(val);
}

GType _g_type_interface()               { return G_TYPE_INTERFACE; }
GType _g_type_char()                    { return G_TYPE_CHAR; }
GType _g_type_uchar()                   { return G_TYPE_UCHAR; }
GType _g_type_boolean()                 { return G_TYPE_BOOLEAN; }
GType _g_type_int()                     { return G_TYPE_INT; }
GType _g_type_uint()                    { return G_TYPE_UINT; }
GType _g_type_long()                    { return G_TYPE_LONG; }
GType _g_type_ulong()                   { return G_TYPE_ULONG; }
GType _g_type_int64()                   { return G_TYPE_INT64; }
GType _g_type_uint64()                  { return G_TYPE_UINT64; }
GType _g_type_enum()                    { return G_TYPE_ENUM; }
GType _g_type_flags()                   { return G_TYPE_FLAGS; }
GType _g_type_float()                   { return G_TYPE_FLOAT; }
GType _g_type_double()                  { return G_TYPE_DOUBLE; }
GType _g_type_string()                  { return G_TYPE_STRING; }
GType _g_type_pointer()                 { return G_TYPE_POINTER; }
GType _g_type_boxed()                   { return G_TYPE_BOXED; }
GType _g_type_param()                   { return G_TYPE_PARAM; }
GType _g_type_object()                  { return G_TYPE_OBJECT; }
GType _g_type_gtype()                   { return G_TYPE_GTYPE; }
GType _g_type_variant()                 { return G_TYPE_VARIANT; }
GType _g_type_param_boolean()           { return G_TYPE_PARAM_BOOLEAN; }
GType _g_type_param_boxed()             { return G_TYPE_PARAM_BOXED; }
GType _g_type_param_char()              { return G_TYPE_PARAM_CHAR; }
GType _g_type_param_double()            { return G_TYPE_PARAM_DOUBLE; }
GType _g_type_param_enum()              { return G_TYPE_PARAM_ENUM; }
GType _g_type_param_flags()             { return G_TYPE_PARAM_FLAGS; }
GType _g_type_param_float()             { return G_TYPE_PARAM_FLOAT; }
GType _g_type_param_gtype()             { return G_TYPE_PARAM_GTYPE; }
GType _g_type_param_int()               { return G_TYPE_PARAM_INT; }
GType _g_type_param_int64()             { return G_TYPE_PARAM_INT64; }
GType _g_type_param_long()              { return G_TYPE_PARAM_LONG; }
GType _g_type_param_object()            { return G_TYPE_PARAM_OBJECT; }
GType _g_type_param_override()          { return G_TYPE_PARAM_OVERRIDE; }
GType _g_type_param_param()             { return G_TYPE_PARAM_PARAM; }
GType _g_type_param_pointer()           { return G_TYPE_PARAM_POINTER; }
GType _g_type_param_string()            { return G_TYPE_PARAM_STRING; }
GType _g_type_param_uchar()             { return G_TYPE_PARAM_UCHAR; }
GType _g_type_param_uint()              { return G_TYPE_PARAM_UINT; }
GType _g_type_param_uint64()            { return G_TYPE_PARAM_UINT64; }
GType _g_type_param_ulong()             { return G_TYPE_PARAM_ULONG; }
GType _g_type_param_unichar()           { return G_TYPE_PARAM_UNICHAR; }
GType _g_type_param_value_array()       { return G_TYPE_PARAM_VALUE_ARRAY; }
GType _g_type_param_variant()           { return G_TYPE_PARAM_VARIANT; }

extern void *gGoInterfaceCopyGo(void *boxed);
extern void gGoInterfaceFreeGo(void *boxed);

//-----------------------------------------------------------------------------

extern void gGoclosureMarshalGo(GGoClosure*, GValue*, int32_t, GValue*);
extern void gGoclosureFinalizeGo(GGoClosure*);

GType _g_type_go_interface()
{
        static GType go_interface_type = G_TYPE_NONE;
        if (go_interface_type == G_TYPE_NONE) {
                go_interface_type = g_boxed_type_register_static("gointerface",
                                                                 gGoInterfaceCopyGo,
                                                                 gGoInterfaceFreeGo);
        }
        return go_interface_type;
}

static void gGoclosureFinalize(void *notify_data, GClosure *closure)
{
        GGoClosure *goclosure = (GGoClosure*)closure;
        gGoclosureFinalizeGo(goclosure);
}

static void gGoclosureMarshal(GClosure *closure, GValue *return_value,
                                uint32_t n_param_values, const GValue *param_values,
                                void *invocation_hint, void *data)
{
        gGoclosureMarshalGo((GGoClosure*)closure,
                               return_value,
                               n_param_values,
                               (GValue*)param_values);
}



void *gGoclosureGetFunc(GGoClosure *clo) {
        return clo->func;
}

void *gGoclosureGetRecv(GGoClosure *clo) {
        return clo->recv;
}

static void _gtk_init(void *argc, void *argv)
{
        gtk_init((int *)argc, (char ***)argv);
}

GGoClosure *gGoclosureNew(void *func, void *recv)
{
        GClosure *closure;
        GGoClosure *goclosure;

        closure = g_closure_new_simple(sizeof(GGoClosure), 0);
        goclosure = (GGoClosure*)closure;
        memset(goclosure->func, 0, sizeof(goclosure->func));
        memset(goclosure->recv, 0, sizeof(goclosure->recv));
        if (func)
                memcpy(goclosure->func, func, sizeof(goclosure->func));
        if (recv)
                memcpy(goclosure->recv, recv, sizeof(goclosure->recv));

        g_closure_add_finalize_notifier(closure, 0, gGoclosureFinalize);
        g_closure_set_marshal(closure, gGoclosureMarshal);

        return goclosure;
}

*/
import "C"

import (
        "unsafe"
)

type WindowType C.int

type PositionType C.int

type ResizeMode C.int

type Window struct {
        c *C.GtkWindow
}

type Widget struct {
        c *C.GtkWidget
}

type Button struct {
        c *C.GtkButton
}

type Container struct {
        c *C.GtkContainer
}

type VScrollbar struct {
        c *C.GtkVScrollbar
}

type HandleBox struct {
        c *C.GtkHandleBox
}

type HPaned struct {
        c *C.GtkHPaned
}
type HScale struct {
        c *C.GtkHScale
}
type ColorSelection struct {
        c *C.GtkColorSelection
}
type HButtonBox struct {
        c *C.GtkHButtonBox
}
type VBox struct {
        c *C.GtkVBox
}
type TearoffMenuItem struct {
        c *C.GtkTearoffMenuItem
}
type RcStyle struct {
        c *C.GtkRcStyle
}
type HBox struct {
        c *C.GtkHBox
}

type HSeparator struct {
        c *C.GtkHSeparator
}
type Table struct {
        c *C.GtkTable
}
type VScale struct {
        c *C.GtkVScale
}
type VSeparator struct {
        c *C.GtkVSeparator
}
type FontSelection struct {
        c *C.GtkFontSelection
}
type FontSelectionDialog struct {
        c *C.GtkFontSelectionDialog
}
type HSV struct {
        c *C.GtkHSV
}
type Style struct {
        c *C.GtkStyle
}
type Gradient struct {
        c *C.GtkGradient
}
type ColorSelectionDialog struct {
        c *C.GtkColorSelectionDialog
}

type HScrollbar struct {
        c *C.GtkHScrollbar
}
type VButtonBox struct {
        c *C.GtkVButtonBox
}
type VPaned struct {
        c *C.GtkVPaned
}
type SymbolicColor struct {
        c *C.GtkSymbolicColor
}
type TextCellAccessible struct {
        c *C.GtkTextCellAccessible
}
type ContainerAccessible struct {
        c *C.GtkContainerAccessible
}
type ButtonAccessible struct {
        c *C.GtkButtonAccessible
}
type CellAccessible struct {
        c *C.GtkCellAccessible
}
type TreeViewAccessible struct {
        c *C.GtkTreeViewAccessible
}
type RangeAccessible struct {
        c *C.GtkRangeAccessible
}
type NotebookPageAccessible struct {
        c *C.GtkNotebookPageAccessible
}
type ArrowAccessible struct {
        c *C.GtkArrowAccessible
}
type ComboBoxAccessible struct {
        c *C.GtkComboBoxAccessible
}
type ExpanderAccessible struct {
        c *C.GtkExpanderAccessible
}
type TextViewAccessible struct {
        c *C.GtkTextViewAccessible
}
type BooleanCellAccessible struct {
        c *C.GtkBooleanCellAccessible
}
type MenuItemAccessible struct {
        c *C.GtkMenuItemAccessible
}
type WidgetAccessible struct {
        c *C.GtkWidgetAccessible
}
type ScaleButtonAccessible struct {
        c *C.GtkScaleButtonAccessible
}
type ImageAccessible struct {
        c *C.GtkImageAccessible
}
type PanedAccessible struct {
        c *C.GtkPanedAccessible
}
type RadioButtonAccessible struct {
        c *C.GtkRadioButtonAccessible
}
type RendererCellAccessible struct {
        c *C.GtkRendererCellAccessible
}
type ProgressBarAccessible struct {
        c *C.GtkProgressBarAccessible
}
type CheckMenuItemAccessible struct {
        c *C.GtkCheckMenuItemAccessible
}
type IconViewAccessible struct {
        c *C.GtkIconViewAccessible
}
type RadioMenuItemAccessible struct {
        c *C.GtkRadioMenuItemAccessible
}
type LockButtonAccessible struct {
        c *C.GtkLockButtonAccessible
}
type ContainerCellAccessible struct {
        c *C.GtkContainerCellAccessible
}
type SwitchAccessible struct {
        c *C.GtkSwitchAccessible
}
type StatusbarAccessible struct {
        c *C.GtkStatusbarAccessible
}
type SpinButtonAccessible struct {
        c *C.GtkSpinButtonAccessible
}
type FrameAccessible struct {
        c *C.GtkFrameAccessible
}
type LevelBarAccessible struct {
        c *C.GtkLevelBarAccessible
}
type EntryAccessible struct {
        c *C.GtkEntryAccessible
}
type MenuAccessible struct {
        c *C.GtkMenuAccessible
}
type ToplevelAccessible struct {
        c *C.GtkToplevelAccessible
}
type MenuShellAccessible struct {
        c *C.GtkMenuShellAccessible
}
type WindowAccessible struct {
        c *C.GtkWindowAccessible
}
type LabelAccessible struct {
        c *C.GtkLabelAccessible
}
type ScaleAccessible struct {
        c *C.GtkScaleAccessible
}
type ToggleButtonAccessible struct {
        c *C.GtkToggleButtonAccessible
}
type CellAccessibleParent struct {
        c *C.GtkCellAccessibleParent
}
type NotebookAccessible struct {
        c *C.GtkNotebookAccessible
}
type SpinnerAccessible struct {
        c *C.GtkSpinnerAccessible
}
type ImageCellAccessible struct {
        c *C.GtkImageCellAccessible
}
type ScrolledWindowAccessible struct {
        c *C.GtkScrolledWindowAccessible
}
type LinkButtonAccessible struct {
        c *C.GtkLinkButtonAccessible
}

type Grid struct {
        c *C.GtkGrid
}

type WindowGroup struct {
        c *C.GtkWindowGroup
}

type GList struct {
        c *C.GList
}

const (
        WINDOW_TOPLEVEL WindowType = C.GTK_WINDOW_TOPLEVEL
        WINDOW_POPUP    WindowType = C.GTK_WINDOW_POPUP
)

const (
        POS_LEFT   PositionType = C.GTK_POS_LEFT
        POS_RIGHT  PositionType = C.GTK_POS_RIGHT
        POS_TOP    PositionType = C.GTK_POS_TOP
        POS_BOTTOM PositionType = C.GTK_POS_BOTTOM
)

const (
        RESIZE_PARENT    ResizeMode = C.GTK_RESIZE_PARENT
        RESIZE_QUEUE     ResizeMode = C.GTK_RESIZE_QUEUE
        RESIZE_IMMEDIATE ResizeMode = C.GTK_RESIZE_IMMEDIATE
)

func _GString(s string) *C.gchar {
        return (*C.gchar)(C.CString(s))
}

func _GFree(v unsafe.Pointer) {
        C.g_free(C.gpointer(v))
}

func Init(args []string) {
        argc := C.int(len(args))
        iLen := len(args)
        if iLen != 0 {
                argv := make([]*C.char, iLen)
                for i := 0; i < iLen; i++ {
                        argv[i] = C.CString(args[i])
                }
                C._gtk_init(unsafe.Pointer(&argc), unsafe.Pointer(&argv))
                for i := 0; i < iLen; i++ {
                        args[i] = C.GoString(argv[i])
                        C.free(unsafe.Pointer(argv[i]))
                }
        } else {
                C._gtk_init(nil, nil)
        }
}

func Main() {
        C.gtk_main()
}

func MainQuit() {
        C.gtk_main_quit()
}
