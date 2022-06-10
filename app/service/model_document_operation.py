# -- coding: utf-8 --
from bs4 import BeautifulSoup
from config.settings import IMAGE_FORMAT
from config.omc import omc
import base64

def GetModelDocument(model_name):
    doc_data = omc.getDocumentationAnnotation(model_name)
    if doc_data[0]:
        html = doc_data[0]
        soup = BeautifulSoup(html, features="html.parser")
        image_all = soup.find_all('img')
        a_href_all = soup.find_all('a')
        for image in image_all:
            src = image['src']
            image_suffix = src.split(".")[-1]
            if image_suffix in IMAGE_FORMAT:
                image_file = omc.uriToFilename(image['src'])
                with open(image_file, "rb") as f:
                    image_base64 = base64.b64encode(f.read())
                    image['src'] = "data:image/jpeg;base64," + image_base64.decode()
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
