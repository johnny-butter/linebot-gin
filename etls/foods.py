import csv
import os
import requests
import zipfile


print('-----> Download file(zip) => unzip the file')
DATA_URL = 'https://data.fda.gov.tw/opendata/exportDataList.do?method=ExportData&InfoId=20&logType=2'

DATA_FILE_NAME_ZIP = 'foods.zip'
DATA_FILE_NAME = '20_2.csv'

resp = requests.get(DATA_URL)
with open(DATA_FILE_NAME_ZIP, 'wb') as f:
    f.write(resp.content)

with zipfile.ZipFile(DATA_FILE_NAME_ZIP, 'r') as zip_f:
    zip_f.extractall('./')


print('-----> Read data source and turn it to dictionary')
out_dict = {}

ingredients = ['熱量', '飽和脂肪', '反式脂肪', '膽固醇', '鋅', '鐵']

with open(DATA_FILE_NAME, 'r', encoding='utf-8') as f:
    reader = csv.reader(f)
    next(reader)

    for rows in reader:
        if not out_dict.get(rows[3]):
            out_dict[rows[3]] = {}
            out_dict[rows[3]]['amount_per_100g'] = {}

            out_dict[rows[3]]['name'] = rows[3]
            out_dict[rows[3]]['common_names'] = rows[4]
            out_dict[rows[3]]['eng_name'] = rows[5]
            out_dict[rows[3]]['category'] = rows[0]
            out_dict[rows[3]]['code'] = rows[2]

        try:
            if rows[9] not in ingredients:
                continue

            amount_val = int(float(rows[11]))
            if amount_val == 0:
                continue

            out_dict[rows[3]]['amount_per_100g'][rows[9]] = f'{amount_val} {rows[10]}'
        except Exception as e:
            print(str(e), rows)


print('-----> Output csv files from dictionary')
OUT_FILE_NAME = 'foods.csv'
OUT_FILE_NAME_2 = 'food_ingredients.csv'

OUT_LABELS = ['id', 'name', 'eng_name', 'category', 'common_names', 'code']
OUT_LABELS_2 = ['id', 'name', 'amount', 'food_id']

OUT_ID = 1
OUT_ID_2 = 1

try:
    with open(OUT_FILE_NAME, 'w', encoding='utf-8', newline='') as f:
        with open(OUT_FILE_NAME_2, 'w', encoding='utf-8', newline='') as f2:
            writer = csv.DictWriter(f, fieldnames=OUT_LABELS)
            writer.writeheader()

            writer2 = csv.writer(f2)
            writer2.writerow(OUT_LABELS_2)

            for _, elem in out_dict.items():
                amount_elem = elem.pop('amount_per_100g')

                elem['id'] = OUT_ID
                writer.writerow(elem)

                for k, v in amount_elem.items():
                    writer2.writerow([OUT_ID_2, k, v, OUT_ID])
                    OUT_ID_2 += 1

                OUT_ID += 1
except Exception as e:
    print(str(e))


print('-----> Remove redundant files')
os.remove(DATA_FILE_NAME_ZIP)
os.remove(DATA_FILE_NAME)


print('-----> Finished')
