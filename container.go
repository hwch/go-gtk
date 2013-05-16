package gtk

/*
#cgo pkg-config: gobject-2.0 gtk+-3.0
#include <glib-object.h>
#include <gtk/gtk.h>
#include "gobj.h"
*/
import "C"

// void    gtk_container_add		 (GtkContainer	   *container,
// 					  GtkWidget	   *widget);
func (c *Container) Add(w *Widget) {
        C.gtk_container_add(c.c, w.c)
}

func (b *Container) CType() *C.GtkContainer {
        return b.c
}

func (c *Container) SetBorderWidth(w uint32) {
        C.gtk_container_set_border_width(c.c, C.guint(w))
}

// guint   gtk_container_get_border_width   (GtkContainer     *container);
func (c *Container) GetBorderWidth() uint32 {
        return uint32(C.gtk_container_get_border_width(c.c))
}

// void    gtk_container_remove		 (GtkContainer	   *container,
// 					  GtkWidget	   *widget);
func (c *Container) Remove(widget Widget) {
        C.gtk_container_remove(c.c, widget.c)
}

// void    gtk_container_set_resize_mode    (GtkContainer     *container,
// 					  GtkResizeMode     resize_mode);
func (c *Container) SetResizeMode(resize_mode ResizeMode) {
        C.gtk_container_set_resize_mode(c.c, C.GtkResizeMode(resize_mode))
}

// GtkResizeMode gtk_container_get_resize_mode (GtkContainer     *container);
func (c *Container) GetResizeMode() ResizeMode {
        return ResizeMode(C.gtk_container_get_resize_mode(c.c))
}

// void    gtk_container_check_resize       (GtkContainer     *container);

// void     gtk_container_foreach      (GtkContainer       *container,
// 				     GtkCallback         callback,
// 				     gpointer            callback_data);
// GList*   gtk_container_get_children     (GtkContainer       *container);

// void     gtk_container_propagate_draw   (GtkContainer   *container,
// 					 GtkWidget      *child,
// 					 cairo_t        *cr);

// void     gtk_container_set_focus_chain  (GtkContainer   *container,
//                                          GList          *focusable_widgets);
// gboolean gtk_container_get_focus_chain  (GtkContainer   *container,
// 					 GList         **focusable_widgets);
// void     gtk_container_unset_focus_chain (GtkContainer  *container);
// void   gtk_container_set_reallocate_redraws (GtkContainer    *container,
// 					     gboolean         needs_redraws);
// void   gtk_container_set_focus_child	   (GtkContainer     *container,
// 					    GtkWidget	     *child);
// GtkWidget *
//        gtk_container_get_focus_child	   (GtkContainer     *container);
// void   gtk_container_set_focus_vadjustment (GtkContainer     *container,
// 					    GtkAdjustment    *adjustment);
// GtkAdjustment *gtk_container_get_focus_vadjustment (GtkContainer *container);
// void   gtk_container_set_focus_hadjustment (GtkContainer     *container,
// 					    GtkAdjustment    *adjustment);
// GtkAdjustment *gtk_container_get_focus_hadjustment (GtkContainer *container);

// void    gtk_container_resize_children      (GtkContainer     *container);

// GType   gtk_container_child_type	   (GtkContainer     *container);

// void         gtk_container_class_install_child_property (GtkContainerClass *cclass,
// 							 guint		    property_id,
// 							 GParamSpec	   *pspec);
// GParamSpec*  gtk_container_class_find_child_property	(GObjectClass	   *cclass,
// 							 const gchar	   *property_name);
// GParamSpec** gtk_container_class_list_child_properties	(GObjectClass	   *cclass,
// 							 guint		   *n_properties);
// void         gtk_container_add_with_properties		(GtkContainer	   *container,
// 							 GtkWidget	   *widget,
// 							 const gchar	   *first_prop_name,
// 							 ...) G_GNUC_NULL_TERMINATED;
// void         gtk_container_child_set			(GtkContainer	   *container,
// 							 GtkWidget	   *child,
// 							 const gchar	   *first_prop_name,
// 							 ...) G_GNUC_NULL_TERMINATED;
// void         gtk_container_child_get			(GtkContainer	   *container,
// 							 GtkWidget	   *child,
// 							 const gchar	   *first_prop_name,
// 							 ...) G_GNUC_NULL_TERMINATED;
// void         gtk_container_child_set_valist		(GtkContainer	   *container,
// 							 GtkWidget	   *child,
// 							 const gchar	   *first_property_name,
// 							 va_list	    var_args);
// void         gtk_container_child_get_valist		(GtkContainer	   *container,
// 							 GtkWidget	   *child,
// 							 const gchar	   *first_property_name,
// 							 va_list	    var_args);
// void	     gtk_container_child_set_property		(GtkContainer	   *container,
// 							 GtkWidget	   *child,
// 							 const gchar	   *property_name,
// 							 const GValue	   *value);
// void	     gtk_container_child_get_property		(GtkContainer	   *container,
// 							 GtkWidget	   *child,
// 							 const gchar	   *property_name,
// 	                                                 GValue		   *value);

// GDK_AVAILABLE_IN_3_2
// void gtk_container_child_notify (GtkContainer *container,
//                                  GtkWidget    *child,
//                                  const gchar  *child_property);
// void    gtk_container_forall		     (GtkContainer *container,
// 					      GtkCallback   callback,
// 					      gpointer	    callback_data);

// void    gtk_container_class_handle_border_width (GtkContainerClass *klass);

// GtkWidgetPath * gtk_container_get_path_for_child (GtkContainer      *container,
//                                                   GtkWidget         *child);
