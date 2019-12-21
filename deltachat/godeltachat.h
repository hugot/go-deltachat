uintptr_t godeltachat_eventhandler(dc_context_t* context, int event,
                                 uintptr_t data1, uintptr_t data2);

dc_context_t* godeltachat_create_context();

int godeltachat_event_data1_is_string(int event);

int godeltachat_event_data2_is_string(int event);
