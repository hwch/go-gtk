package gtk

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include <gtk/gtk-a11y.h>

*/
import "C"
import "unsafe"

func ToWindow(ip unsafe.Pointer) *Window {
        tp := C.gtk_window_get_type()
        win := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Window{(*C.GtkWindow)(unsafe.Pointer(win))}
}

func ToWidget(ip unsafe.Pointer) *Widget {
        tp := C.gtk_widget_get_type()
        wd := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Widget{(*C.GtkWidget)(unsafe.Pointer(wd))}
}

func ToButton(ip unsafe.Pointer) *Button {
        tp := C.gtk_button_get_type()
        bt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Button{(*C.GtkButton)(unsafe.Pointer(bt))}
}

func ToVscrollbar(ip unsafe.Pointer) *VScrollbar {
        tp := C.gtk_vscrollbar_get_type()
        vs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &VScrollbar{(*C.GtkVScrollbar)(unsafe.Pointer(vs))}
}

func ToHandleBox(ip unsafe.Pointer) *HandleBox {
        tp := C.gtk_handle_box_get_type()
        hb := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HandleBox{(*C.GtkHandleBox)(unsafe.Pointer(hb))}
}

func ToHpaned(ip unsafe.Pointer) *HPaned {
        tp := C.gtk_hpaned_get_type()
        hp := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HPaned{(*C.GtkHPaned)(unsafe.Pointer(hp))}
}

func ToHscale(ip unsafe.Pointer) *HScale {
        tp := C.gtk_hscale_get_type()
        hs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HScale{(*C.GtkHScale)(unsafe.Pointer(hs))}
}

func ToColorSelection(ip unsafe.Pointer) *ColorSelection {
        tp := C.gtk_color_selection_get_type()
        cs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ColorSelection{(*C.GtkColorSelection)(unsafe.Pointer(cs))}
}

func ToHButtonBox(ip unsafe.Pointer) *HButtonBox {
        tp := C.gtk_hbutton_box_get_type()
        hbb := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HButtonBox{(*C.GtkHButtonBox)(unsafe.Pointer(hbb))}
}

func ToVBox(ip unsafe.Pointer) *VBox {
        tp := C.gtk_vbox_get_type()
        vb := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &VBox{(*C.GtkVBox)(unsafe.Pointer(vb))}
}

func ToTearoffMenuItem(ip unsafe.Pointer) *TearoffMenuItem {
        tp := C.gtk_tearoff_menu_item_get_type()
        tmi := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &TearoffMenuItem{(*C.GtkTearoffMenuItem)(unsafe.Pointer(tmi))}
}

func ToRcStyle(ip unsafe.Pointer) *RcStyle {
        tp := C.gtk_rc_style_get_type()
        rs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &RcStyle{(*C.GtkRcStyle)(unsafe.Pointer(rs))}
}

func ToHBox(ip unsafe.Pointer) *HBox {
        tp := C.gtk_hbox_get_type()
        hb := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HBox{(*C.GtkHBox)(unsafe.Pointer(hb))}
}

func ToHSeparator(ip unsafe.Pointer) *HSeparator {
        tp := C.gtk_hseparator_get_type()
        hs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HSeparator{(*C.GtkHSeparator)(unsafe.Pointer(hs))}
}

func ToTable(ip unsafe.Pointer) *Table {
        tp := C.gtk_table_get_type()
        t := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Table{(*C.GtkTable)(unsafe.Pointer(t))}
}

func ToVScale(ip unsafe.Pointer) *VScale {
        tp := C.gtk_vscale_get_type()
        vs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &VScale{(*C.GtkVScale)(unsafe.Pointer(vs))}
}

func ToVSeparator(ip unsafe.Pointer) *VSeparator {
        tp := C.gtk_vseparator_get_type()
        vb := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &VSeparator{(*C.GtkVSeparator)(unsafe.Pointer(vb))}
}

func ToFontSelection(ip unsafe.Pointer) *FontSelection {
        tp := C.gtk_font_selection_get_type()
        fs := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &FontSelection{(*C.GtkFontSelection)(unsafe.Pointer(fs))}
}

func ToFontSelectionDialog(ip unsafe.Pointer) *FontSelectionDialog {
        tp := C.gtk_font_selection_dialog_get_type()
        fsd := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &FontSelectionDialog{(*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd))}
}

func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        hsv := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HSV{(*C.GtkHSV)(unsafe.Pointer(hsv))}
}

func ToStyle(ip unsafe.Pointer) *Style {
        tp := C.gtk_style_get_type()
        s := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Style{(*C.GtkStyle)(unsafe.Pointer(s))}
}

func ToGradient(ip unsafe.Pointer) *Gradient {
        tp := C.gtk_gradient_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Gradient{(*C.GtkGradient)(unsafe.Pointer(gt))}
}

func ToColorSelectionDialog(ip unsafe.Pointer) *ColorSelectionDialog {
        tp := C.gtk_color_selection_dialog_get_type()
        csd := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ColorSelectionDialog{(*C.GtkColorSelectionDialog)(unsafe.Pointer(csd))}
}

func ToHScrollbar(ip unsafe.Pointer) *HScrollbar {
        tp := C.gtk_hscrollbar_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &HScrollbar{(*C.GtkHScrollbar)(unsafe.Pointer(gt))}
}

func ToVButtonBox(ip unsafe.Pointer) *VButtonBox {
        tp := C.gtk_vbutton_box_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &VButtonBox{(*C.GtkVButtonBox)(unsafe.Pointer(gt))}
}

func ToVPaned(ip unsafe.Pointer) *VPaned {
        tp := C.gtk_vpaned_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &VPaned{(*C.GtkVPaned)(unsafe.Pointer(gt))}
}

func ToSymbolicColor(ip unsafe.Pointer) *SymbolicColor {
        tp := C.gtk_symbolic_color_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &SymbolicColor{(*C.GtkSymbolicColor)(unsafe.Pointer(gt))}
}

func ToTextCellAccessible(ip unsafe.Pointer) *TextCellAccessible {
        tp := C.gtk_text_cell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &TextCellAccessible{(*C.GtkTextCellAccessible)(unsafe.Pointer(gt))}
}

func ToContainerAccessible(ip unsafe.Pointer) *ContainerAccessible {
        tp := C.gtk_container_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ContainerAccessible{(*C.GtkContainerAccessible)(unsafe.Pointer(gt))}
}

func ToButtonAccessible(ip unsafe.Pointer) *ButtonAccessible {
        tp := C.gtk_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ButtonAccessible{(*C.GtkButtonAccessible)(unsafe.Pointer(gt))}
}

func ToCellAccessible(ip unsafe.Pointer) *CellAccessible {
        tp := C.gtk_cell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &CellAccessible{(*C.GtkCellAccessible)(unsafe.Pointer(gt))}
}

func ToTreeViewAccessible(ip unsafe.Pointer) *TreeViewAccessible {
        tp := C.gtk_tree_view_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &TreeViewAccessible{(*C.GtkTreeViewAccessible)(unsafe.Pointer(gt))}
}

func ToRangeAccessible(ip unsafe.Pointer) *RangeAccessible {
        tp := C.gtk_range_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &RangeAccessible{(*C.GtkRangeAccessible)(unsafe.Pointer(gt))}
}

func ToNotebookPageAccessible(ip unsafe.Pointer) *NotebookPageAccessible {
        tp := C.gtk_notebook_page_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &NotebookPageAccessible{(*C.GtkNotebookPageAccessible)(unsafe.Pointer(gt))}
}

func ToArrowAccessible(ip unsafe.Pointer) *ArrowAccessible {
        tp := C.gtk_arrow_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ArrowAccessible{(*C.GtkArrowAccessible)(unsafe.Pointer(gt))}
}

func ToComboBoxAccessible(ip unsafe.Pointer) *ComboBoxAccessible {
        tp := C.gtk_combo_box_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ComboBoxAccessible{(*C.GtkComboBoxAccessible)(unsafe.Pointer(gt))}
}

func ToExpanderAccessible(ip unsafe.Pointer) *ExpanderAccessible {
        tp := C.gtk_expander_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ExpanderAccessible{(*C.GtkExpanderAccessible)(unsafe.Pointer(gt))}
}

func ToTextViewAccessible(ip unsafe.Pointer) *TextViewAccessible {
        tp := C.gtk_text_view_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &TextViewAccessible{(*C.GtkTextViewAccessible)(unsafe.Pointer(gt))}
}

func ToBooleanCellAccessible(ip unsafe.Pointer) *BooleanCellAccessible {
        tp := C.gtk_boolean_cell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &BooleanCellAccessible{(*C.GtkBooleanCellAccessible)(unsafe.Pointer(gt))}
}

func ToMenuItemAccessible(ip unsafe.Pointer) *MenuItemAccessible {
        tp := C.gtk_menu_item_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &MenuItemAccessible{(*C.GtkMenuItemAccessible)(unsafe.Pointer(gt))}
}

func ToWidgetAccessible(ip unsafe.Pointer) *WidgetAccessible {
        tp := C.gtk_widget_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &WidgetAccessible{(*C.GtkWidgetAccessible)(unsafe.Pointer(gt))}
}

func ToScaleButtonAccessible(ip unsafe.Pointer) *ScaleButtonAccessible {
        tp := C.gtk_scale_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ScaleButtonAccessible{(*C.GtkScaleButtonAccessible)(unsafe.Pointer(gt))}
}

func ToImageAccessible(ip unsafe.Pointer) *ImageAccessible {
        tp := C.gtk_image_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ImageAccessible{(*C.GtkImageAccessible)(unsafe.Pointer(gt))}
}

func ToPanedAccessible(ip unsafe.Pointer) *PanedAccessible {
        tp := C.gtk_paned_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &PanedAccessible{(*C.GtkPanedAccessible)(unsafe.Pointer(gt))}
}

func ToRadioButtonAccessible(ip unsafe.Pointer) *RadioButtonAccessible {
        tp := C.gtk_radio_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &RadioButtonAccessible{(*C.GtkRadioButtonAccessible)(unsafe.Pointer(gt))}
}

func ToRendererCellAccessible(ip unsafe.Pointer) *RendererCellAccessible {
        tp := C.gtk_renderer_cell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &RendererCellAccessible{(*C.GtkRendererCellAccessible)(unsafe.Pointer(gt))}
}

func ToProgressBarAccessible(ip unsafe.Pointer) *ProgressBarAccessible {
        tp := C.gtk_progress_bar_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ProgressBarAccessible{(*C.GtkProgressBarAccessible)(unsafe.Pointer(gt))}
}

func ToCheckMenuItemAccessible(ip unsafe.Pointer) *CheckMenuItemAccessible {
        tp := C.gtk_check_menu_item_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &CheckMenuItemAccessible{(*C.GtkCheckMenuItemAccessible)(unsafe.Pointer(gt))}
}

func ToIconViewAccessible(ip unsafe.Pointer) *IconViewAccessible {
        tp := C.gtk_icon_view_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &IconViewAccessible{(*C.GtkIconViewAccessible)(unsafe.Pointer(gt))}
}

func ToRadioMenuItemAccessible(ip unsafe.Pointer) *RadioMenuItemAccessible {
        tp := C.gtk_radio_menu_item_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &RadioMenuItemAccessible{(*C.GtkRadioMenuItemAccessible)(unsafe.Pointer(gt))}
}

func ToLockButtonAccessible(ip unsafe.Pointer) *LockButtonAccessible {
        tp := C.gtk_lock_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &LockButtonAccessible{(*C.GtkLockButtonAccessible)(unsafe.Pointer(gt))}
}

func ToContainerCellAccessible(ip unsafe.Pointer) *ContainerCellAccessible {
        tp := C.gtk_container_cell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ContainerCellAccessible{(*C.GtkContainerCellAccessible)(unsafe.Pointer(gt))}
}

func ToSwitchAccessible(ip unsafe.Pointer) *SwitchAccessible {
        tp := C.gtk_switch_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &SwitchAccessible{(*C.GtkSwitchAccessible)(unsafe.Pointer(gt))}
}

func ToStatusbarAccessible(ip unsafe.Pointer) *StatusbarAccessible {
        tp := C.gtk_statusbar_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &StatusbarAccessible{(*C.GtkStatusbarAccessible)(unsafe.Pointer(gt))}
}

func ToSpinButtonAccessible(ip unsafe.Pointer) *SpinButtonAccessible {
        tp := C.gtk_spin_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &SpinButtonAccessible{(*C.GtkSpinButtonAccessible)(unsafe.Pointer(gt))}
}

func ToFrameAccessible(ip unsafe.Pointer) *FrameAccessible {
        tp := C.gtk_frame_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &FrameAccessible{(*C.GtkFrameAccessible)(unsafe.Pointer(gt))}
}

func ToLevelBarAccessible(ip unsafe.Pointer) *LevelBarAccessible {
        tp := C.gtk_level_bar_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &LevelBarAccessible{(*C.GtkLevelBarAccessible)(unsafe.Pointer(gt))}
}

func ToEntryAccessible(ip unsafe.Pointer) *EntryAccessible {
        tp := C.gtk_entry_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &EntryAccessible{(*C.GtkEntryAccessible)(unsafe.Pointer(gt))}
}

func ToMenuAccessible(ip unsafe.Pointer) *MenuAccessible {
        tp := C.gtk_menu_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &MenuAccessible{(*C.GtkMenuAccessible)(unsafe.Pointer(gt))}
}

func ToToplevelAccessible(ip unsafe.Pointer) *ToplevelAccessible {
        tp := C.gtk_toplevel_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ToplevelAccessible{(*C.GtkToplevelAccessible)(unsafe.Pointer(gt))}
}

func ToMenuShellAccessible(ip unsafe.Pointer) *MenuShellAccessible {
        tp := C.gtk_menu_shell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &MenuShellAccessible{(*C.GtkMenuShellAccessible)(unsafe.Pointer(gt))}
}

func ToWindowAccessible(ip unsafe.Pointer) *WindowAccessible {
        tp := C.gtk_window_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &WindowAccessible{(*C.GtkWindowAccessible)(unsafe.Pointer(gt))}
}

func ToLabelAccessible(ip unsafe.Pointer) *LabelAccessible {
        tp := C.gtk_label_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &LabelAccessible{(*C.GtkLabelAccessible)(unsafe.Pointer(gt))}
}

func ToScaleAccessible(ip unsafe.Pointer) *ScaleAccessible {
        tp := C.gtk_scale_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ScaleAccessible{(*C.GtkScaleAccessible)(unsafe.Pointer(gt))}
}

func ToToggleButtonAccessible(ip unsafe.Pointer) *ToggleButtonAccessible {
        tp := C.gtk_toggle_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ToggleButtonAccessible{(*C.GtkToggleButtonAccessible)(unsafe.Pointer(gt))}
}

func ToCellAccessibleParent(ip unsafe.Pointer) *CellAccessibleParent {
        tp := C.gtk_cell_accessible_parent_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &CellAccessibleParent{(*C.GtkCellAccessibleParent)(unsafe.Pointer(gt))}
}

func ToNotebookAccessible(ip unsafe.Pointer) *NotebookAccessible {
        tp := C.gtk_notebook_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &NotebookAccessible{(*C.GtkNotebookAccessible)(unsafe.Pointer(gt))}
}

func ToSpinnerAccessible(ip unsafe.Pointer) *SpinnerAccessible {
        tp := C.gtk_spinner_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &SpinnerAccessible{(*C.GtkSpinnerAccessible)(unsafe.Pointer(gt))}
}

func ToImageCellAccessible(ip unsafe.Pointer) *ImageCellAccessible {
        tp := C.gtk_image_cell_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ImageCellAccessible{(*C.GtkImageCellAccessible)(unsafe.Pointer(gt))}
}

func ToScrolledWindowAccessible(ip unsafe.Pointer) *ScrolledWindowAccessible {
        tp := C.gtk_scrolled_window_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &ScrolledWindowAccessible{(*C.GtkScrolledWindowAccessible)(unsafe.Pointer(gt))}
}

func ToLinkButtonAccessible(ip unsafe.Pointer) *LinkButtonAccessible {
        tp := C.gtk_link_button_accessible_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &LinkButtonAccessible{(*C.GtkLinkButtonAccessible)(unsafe.Pointer(gt))}
}

func ToContainer(ip unsafe.Pointer) *Container {
        tp := C.gtk_container_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Container{(*C.GtkContainer)(unsafe.Pointer(gt))}
}

func ToGrid(ip unsafe.Pointer) *Grid {
        tp := C.gtk_grid_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &Grid{(*C.GtkGrid)(unsafe.Pointer(gt))}
}

func ToWindowGroup(ip unsafe.Pointer) *WindowGroup {
        tp := C.gtk_window_group_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return &WindowGroup{(*C.GtkWindowGroup)(unsafe.Pointer(gt))}
}

/*


func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *HSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
func ToHSV(ip unsafe.Pointer) *C.GtkHSV {
        tp := C.gtk_hsv_get_type()
        gt := C.g_type_check_instance_cast((*C.GTypeInstance)(ip), tp)
        return (*C.GtkHSV)(unsafe.Pointer(gt))
}
*/
