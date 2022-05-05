# -- coding: utf-8 --
from bs4 import BeautifulSoup
from config.settings import IMAGE_FORMAT
from config.omc import omc
from library.HW_OBS_operation import HWOBS

def GetModelDocument(model_name):
    doc_data = omc.getDocumentationAnnotation(model_name)
    if doc_data[0]:
        html = doc_data[0]
        soup = BeautifulSoup(html)
        image_all = soup.find_all('img')
        a_href_all = soup.find_all('a')
        for image in image_all:
            src = image['src']
            file_name = src.split('/')[-1]
            image_suffix = src.split(".")[-1]
            if image_suffix in IMAGE_FORMAT:
                image_path = omc.uriToFilename(image['src'])
                obs = HWOBS()
                new_path = "document/images/" + file_name
                HW_res = obs.putFile(new_path, image_path)
                if HW_res["status"] == 200:
                    image['src'] = [HW_res["body"]["objectUrl"]]
                else:
                    return "No document, please check the model name"
            else:
                del src
        for a in a_href_all:
            del a['href']
        return str(soup)
    else:
        return "No document"


def SetModelDocument(model_name, doc_data):
    result = omc.setDocumentationAnnotation(model_name, doc_data)
    if result:
        return True
    else:
        return False
