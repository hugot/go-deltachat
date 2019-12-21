#include <deltachat.h>
#include <stdio.h>

extern uintptr_t godeltachat_eventhandler_proxy(dc_context_t* context, int event,
                             uintptr_t data1, uintptr_t data2);

uintptr_t godeltachat_eventhandler(dc_context_t* context, int event,
                             uintptr_t data1, uintptr_t data2)
{
  return godeltachat_eventhandler_proxy(context, event, data1, data2);
}

// Context creation because passing a C function as callback value from go does not seeem
// to work
dc_context_t* godeltachat_create_context()
{
  return dc_context_new(godeltachat_eventhandler, NULL, NULL);
}

// Macro wrappers because cgo does not support calling macros
int godeltachat_event_data1_is_string(int event) {
  return DC_EVENT_DATA1_IS_STRING(event);
}

int godeltachat_event_data2_is_string(int event)
{
  return DC_EVENT_DATA2_IS_STRING(event);
}
