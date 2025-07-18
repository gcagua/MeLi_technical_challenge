import sys
import requests

API_ENDPOINT = "https://jsonmock.hackerrank.com/api/tvseries"
def best_in_genre(genre: str) -> str:
    """
    Returns the TV series with the highest IMDB rating for a given genre
    If two series have the same rating, the smaller name is returned
    Python version is 3.13.5
    """
        
    try:
        validated_genre = validate_genre(genre)
    except Exception as e:
        print(f"An error has occured with the genre:", e)
        return None

    try:
        total_pages = fetch_total_pages()
    except requests.RequestException as e:
        print(f"An error occured while fetching data from the api:", e)
        return None

    return find_best_series_by_genre(validated_genre, total_pages)

def find_best_series_by_genre(genre: str, total_pages: int) -> str:
    """
    Iterates through all pages and returns the name of the best-rated series in the genre.
    """
    max_imbd_rating = -1 * sys.maxsize # initially sets the maximum value with the min value possible, so that when finds a greater value, changes the variable
    series_with_max_imdb_rating = None # aslo sets the series name with maximum value to a non-existent-value so that when finds a greater value, changes the variable

    for page_num in range(1,  total_pages + 1):
        try:
            series = fetch_page_data(page_num)

            for serie in series:
                if not is_valid_entry_series(serie):
                    continue

                try:
                    imdb_rating = parse_imdb_rating(serie["imdb_rating"])
                except Exception as e:
                    print("Imdb rating could not be parsed to float", e)
                    continue
                
                if genre_matches(serie["genre"],genre) and (imdb_rating > max_imbd_rating or (imdb_rating == max_imbd_rating and serie["name"] < series_with_max_imdb_rating)):
                    # validates if the genre is in the list of genres and if the current value is greater than the max_value until now
                    # also validates if the current value is equal to the max_value until now but the string is lesser 
                    max_imbd_rating = imdb_rating
                    series_with_max_imdb_rating = serie["name"]

        except Exception as e:
            print(f"An error occured while fetching data from the pag {page_num}:", e)
            continue
    return series_with_max_imdb_rating

def fetch_page_data(page_num: int) -> list:
    """
    Fetches a specific page of series data from the API.
    """
    try:
        response = requests.get(f"{API_ENDPOINT}?page={page_num}", 
                                verify=True, 
                                timeout=30
                            )
        response.raise_for_status()
        page_data = response.json()
        data: list = page_data.get("data")
        return data
    except Exception as e:
        raise ValueError("API response has no `data` object")

def is_valid_entry_series(serie:dict) -> bool:
    """
    Validates that the serie entry contains required fields.
    """
    return isinstance(serie, dict) and "genre" in serie and "imdb_rating" in serie and "name" in serie

def genre_matches (serie_genres: str, target_genre: str) -> bool:
    """
    Checks if the target genre is present in the serie's genres.
    """
    genre_list: list[str] = [g.strip() for g in serie_genres.lower().split(",")]
    return target_genre in genre_list

def parse_imdb_rating(value) -> float:
    """
    Parses IMDB rating to float
    """
    try:
        return float(value)
    except Exception as e:
        raise ValueError("Imdb rating could not be parsed to float", e)

def validate_genre (genre: str) -> str:
    """
    Validates and formats the genre input.
    """
    if not isinstance(genre, str) or len(genre) == 0:
        raise ValueError("Genre must be a non-empty string")
    return genre.lower().strip()

def fetch_total_pages() -> int :
    """
    Sends a request to the API to retrieve the total number of pages.
    """
    response = requests.get(
        API_ENDPOINT,
        verify = True,
        timeout = 30
    )
    response.raise_for_status()
    data = response.json()
        
    if "total_pages" not in data:
        raise ValueError("Property total pages is missing from response")
    total_pages = data["total_pages"]
    return total_pages

if __name__ == "__main__":
    print(best_in_genre("")) # Genre must be a non empty string
    print(best_in_genre(2)) # Genre must be a non empty string

    print(best_in_genre("akhdasd")) # Genre does not exist, result should be none
    print(best_in_genre("animation")) # Genre exists, two series have the same score (Rick and morty & Avatar: the Last airbender)
    print(best_in_genre("action ")) # Genre exists and has only one series with the greatest score: GoT
