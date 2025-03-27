package errorcode

const (
	ErrorUserValidate          = "error_user_validate"
	ErrorUserBcryptError       = "error_user_bcrypt_error"
	ErrorAvatarFileError       = "error_avatar_file_error"
	ErrorUserDecodeError       = "error_user_decode_error"
	ErrorUserIdNotProvided     = "error_user_id_not_provided"
	ErrorSaveUserError         = "error_save_user_error"
	ErrorFindByEmailError      = "error_find_by_email_error"
	ErrorFindByIdError         = "error_find_by_id_error"
	ErrorUpdateUserError       = "error_update_user_error"
	ErrorSaveOrderError        = "error_save_order_error"
	ErrorSaveOrderItemError    = "error_save_order_item_error"
	ErrorSendS3File            = "error_send_s3_file_error"
	ErrorDeleteS3File          = "error_delete_s3_file_error"
	ErrorSendSQSMessageMarshal = "error_send_sqs_message_marshal_error"
	ErrorSendSQSMessage        = "error_send_sqs_message_error"
	ErrorFindOrderError        = "error_find_order_error"
	ErrorUpdateOrderError      = "error_update_order_error"
)
