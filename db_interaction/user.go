package db_interaction

func Post() {
	return
}

func Get(name string, discord_id string, props []string) ([]map[string]interface{}, error) {
	var res []map[string]interface{} = []map[string]interface{}{}
	res = append(res, map[string]interface{}{"name": "someName1"})
	res = append(res, map[string]interface{}{"name": "someName2"})
	return res, nil
}
