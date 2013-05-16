#pragma once

typedef struct _GGoClosure GGoClosure;

struct _GGoClosure {
        GClosure closure;
        void *func[2];
        void *recv[2];
};

GType _g_param_spec_type(GParamSpec *pspec);
GType _g_param_spec_value_type(GParamSpec *pspec);
GType _g_object_type(GObject *obj);
GType _g_value_type(GValue *val);
GType _g_type_interface();
GType _g_type_char();
GType _g_type_uchar();
GType _g_type_boolean();
GType _g_type_int();
GType _g_type_uint();
GType _g_type_long();
GType _g_type_ulong();
GType _g_type_int64();
GType _g_type_uint64();
GType _g_type_enum();
GType _g_type_flags();
GType _g_type_float();
GType _g_type_double();
GType _g_type_string();
GType _g_type_pointer();
GType _g_type_boxed();
GType _g_type_param();
GType _g_type_object();
GType _g_type_gtype();
GType _g_type_variant();
GType _g_type_go_interface();
GType _g_type_param_boolean();
GType _g_type_param_boxed();
GType _g_type_param_char();
GType _g_type_param_double();
GType _g_type_param_enum();
GType _g_type_param_flags();
GType _g_type_param_float();
GType _g_type_param_gtype();
GType _g_type_param_int();
GType _g_type_param_int64();
GType _g_type_param_long();
GType _g_type_param_object();
GType _g_type_param_override();
GType _g_type_param_param();
GType _g_type_param_pointer();
GType _g_type_param_string();
GType _g_type_param_uchar();
GType _g_type_param_uint();
GType _g_type_param_uint64();
GType _g_type_param_ulong();
GType _g_type_param_unichar();
GType _g_type_param_value_array();
GType _g_type_param_variant();

// null recv is allowed
GGoClosure *gGoclosureNew(void *func, void *recv);
void *gGoclosureGetFunc(GGoClosure *clo);
void *gGoclosureGetRecv(GGoClosure *clo);
