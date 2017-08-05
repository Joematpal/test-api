package utils

// func encode(r *http.Request, data interface{}) (buf bytes.Buffer, err error) {
// 	err = json.NewEncoder(&buf).Encode(data)
// 	if err != nil {
// 		return buf, err
// 	}
// 	b, err := ioutil.ReadAll(buf)
//
// 	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
// 	return buf, nil
// }
