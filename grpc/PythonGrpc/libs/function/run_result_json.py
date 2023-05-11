import json

file_path = '../../config/run_result_records.json'


def read_json_file():
    with open(file_path, 'r') as f:
        json_data = json.load(f)
    return json_data


def write_json_file(json_data):
    with open(file_path, 'w') as f:
        json.dump(json_data, f)


def add_item_to_json(item_dict):
    json_data = read_json_file()
    json_data.append(item_dict)
    write_json_file(json_data)


def delete_item_from_json(record_id):
    json_data = read_json_file()
    for item in json_data:
        if item['id'] == record_id:
            json_data.remove(item)
    write_json_file(json_data)


def update_json_item(record_id, updated_data):
    json_data = read_json_file()
    for item in json_data:
        if item['id'] == record_id:
            item.update(updated_data)
    write_json_file(json_data)



