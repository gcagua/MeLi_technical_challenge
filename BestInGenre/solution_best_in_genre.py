import sys
import requests

"""
1. ¿Cuantas páginas máximo se deben recorrer? (hay 20 páginas) ✓ hasta que haya null (?)
2. ¿Importan las mayúsculas y minúsculas en el parámetro del genero?
3. ¿Importan las mayúsculas y minúsculas en la comparación de los géneros?
__________________________________________________________________________________
Casos de manejo de error
* Ninguno de los generos de la lista hacen match con el parámetro
"""
API_ENDPOINT = "https://jsonmock.hackerrank.com/api/tvseries"
def best_in_genre(genre: str) -> str:
    try:
        if not isinstance(genre, str) or len(genre) == 0:
            raise ValueError("Genre must be a non-empty string")
        genre = genre.lower().strip()
    except Exception as e:
        print(f"An error has occured:", e)
        return None

    try:
        response = requests.get(
            API_ENDPOINT,
            verify = True,
            timeout = 30
        )
        response.raise_for_status()
        data = response.json()
        
        if "total_pages" not in data:
            raise ValueError("Property total pages is missing")
        totalPages = data["total_pages"]

    except requests.RequestException as e:
        print(f"An error occured while fetching data from the api:", e)
        return None

    max_imbd_rating = -1 * sys.maxsize
    series_with_max_imdb_rating = "none"
    # TODO REVISAR EL NONE

    for pageNum in range(1,  totalPages + 1):
        try:
            page = requests.get(
                f"{API_ENDPOINT}?page={pageNum}",
                verify = True,
                timeout = 30
            )
        
            page.raise_for_status()
            page_data = page.json()
            if "data" not in page_data:
                continue
        
            series = page_data["data"]

            for serie in series:
                if not isinstance(serie, dict):
                    continue

                if "genre" not in serie or "imdb_rating" not in serie or "name" not in serie:
                    continue

                genres: str = serie["genre"]
                genres = genres.lower()
                genres_list: list[str] = str.split(genres,",")

                try:
                    imdb_rating: float = serie["imdb_rating"]
                except Exception as e:
                    print("Imdb rating could not be parsed to float", e)
                    continue
                
                if genre in genres_list and (imdb_rating > max_imbd_rating or (imdb_rating == max_imbd_rating and serie["name"] < series_with_max_imdb_rating)):
                    max_imbd_rating = imdb_rating
                    series_with_max_imdb_rating = serie["name"]

            # for current_serie in range(per_page):
            #     series_at_current_serie = series[current_serie]
            #     genres: str = series_at_current_serie["genre"]
            #     genres = genres.lower()
            #     if genre in genres and (series_at_current_serie["imdb_rating"] > max_imbd_rating or (series_at_current_serie["imdb_rating"] == max_imbd_rating and series_at_current_serie["name"] > series_with_max_imdb_rating)):
            #         max_imbd_rating = page[current_serie]["imdb_rating"]
            #         series_with_max_imdb_rating = page[current_serie]["name"]
        except Exception as e:
            print(f"An error occured while fetching data from the pag {pageNum}:", e)
            continue
    return series_with_max_imdb_rating


print(best_in_genre(""))