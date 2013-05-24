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

// GtkWidget* gtk_dialog_new      (void);
func NewDialog() *Dialog {
        ret := C.gtk_dialog_new()

        if ret != nil {
                return ToDialog(unsafe.Pointer(ret))
        }
        return nil
}

// GtkWidget* gtk_dialog_new_with_buttons (const gchar     *title,
//                                         GtkWindow       *parent,
//                                         GtkDialogFlags   flags,
//                                         const gchar     *first_button_text,
//                                         ...) G_GNUC_NULL_TERMINATED;

// void       gtk_dialog_add_action_widget (GtkDialog   *dialog,
//                                          GtkWidget   *child,
//                                          gint         response_id);
// GtkWidget* gtk_dialog_add_button        (GtkDialog   *dialog,
//                                          const gchar *button_text,
//                                          gint         response_id);
// void       gtk_dialog_add_buttons       (GtkDialog   *dialog,
//                                          const gchar *first_button_text,
//                                          ...) G_GNUC_NULL_TERMINATED;

// void gtk_dialog_set_response_sensitive (GtkDialog *dialog,
//                                         gint       response_id,
//                                         gboolean   setting);
// void gtk_dialog_set_default_response   (GtkDialog *dialog,
//                                         gint       response_id);
// GtkWidget* gtk_dialog_get_widget_for_response (GtkDialog *dialog,
//                                                gint       response_id);
// gint gtk_dialog_get_response_for_widget (GtkDialog *dialog,
//                                          GtkWidget *widget);

// gboolean gtk_alternative_dialog_button_order (GdkScreen *screen);
// void     gtk_dialog_set_alternative_button_order (GtkDialog *dialog,
//                                                   gint       first_response_id,
//                                                   ...);
// void     gtk_dialog_set_alternative_button_order_from_array (GtkDialog *dialog,
//                                                              gint       n_params,
//                                                              gint      *new_order);

// /* Emit response signal */
// void gtk_dialog_response           (GtkDialog *dialog,
//                                     gint       response_id);

// /* Returns response_id */
// gint gtk_dialog_run                (GtkDialog *dialog);

// GtkWidget * gtk_dialog_get_action_area  (GtkDialog *dialog);
// GtkWidget * gtk_dialog_get_content_area (GtkDialog *dialog);
