package utils

func Response(success bool, err error, data interface{}) map[string]interface{} {

	reqMap := map[string]interface{}{
		"success": success,
	}

	if err != nil {
		reqMap["error"] = err
	}

	if data != nil {
		reqMap["data"] = data
	}

	return reqMap

}
