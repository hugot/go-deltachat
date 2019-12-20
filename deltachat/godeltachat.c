#include <deltachat.h>

extern uintptr_t godeltachat_eventhandler_proxy(dc_context_t* context, int event,
                             uintptr_t data1, uintptr_t data2);

// normally you will have to define function or variables
// in another separate C file to avoid the multiple definition
// errors, however, using "static inline" is a nice workaround
// for simple functions like this one.
uintptr_t godeltachat_eventhandler(dc_context_t* context, int event,
                             uintptr_t data1, uintptr_t data2)
{
  return godeltachat_eventhandler_proxy(context, event, data1, data2);
}

dc_context_t* godeltachat_create_context()
{
  return dc_context_new(godeltachat_eventhandler, NULL, NULL);
}
