from typing import List

MAX_MATRIX_SIZE = 50
MINE = 1
MINE_INDICATOR = 9

def minesweeper(matrix: List[List[int]]) -> List[List[int]]:
    """
    Main function to process the matrix and return the minesweeper result.
    Each non-mine cell is replaced by the number of surrounding mines and each mine cell is replaced by a constant.
    Python version is 3.13.5
    """
    try:
        if validate_matrix(matrix): 
            
            rows = len(matrix)
            result = create_empty_matrix(rows)

            for i in range(rows):
                for j in range(rows):
                    result[i][j] = process_cell(i, j, matrix, rows)
            return result
    except ValueError as v:
        print(f"Value error: {v}")
    except TypeError as t:
        print(f"Type error: {t}")

def create_empty_matrix(size: int) -> List[List[int]]:
    """Creates an empty matrix filled with zeros of given size."""
    matrix_response = [[0] * size for _ in range(size)]
    return matrix_response

def process_cell(i: int, j: int, matrix: List[List[int]], size: int) -> int:
    """
    If the cell is a mine, return the mine indicator.
    Otherwise, return the number of neighboring mines.
    """
    if matrix[i][j] == MINE:
        return MINE_INDICATOR
    return count_neighboring_mines(i, j, matrix, size)

def count_neighboring_mines(i: int, j: int, matrix: List[List[int]], rows: int) -> int:
    """
    Counts the number of mines surrounding the cell at (i, j).
    """
    mine_sum = 0
    for k in range(-1, 2):
        for l in range(-1,2):
            if k == 0 and l == 0:
                continue
            if (i + k >= 0 and i + k <= rows -1) and (j + l >= 0 and j + l <= rows -1) and matrix[i+k][j+l] == 1:
                mine_sum += 1
    return mine_sum
    
def validate_matrix(matrix: List[List[int]]) -> bool:
    """
    Validates that the matrix is:
    - Not empty
    - A square list of lists
    - Only contains 0s and 1s
    - Does not exceed MAX_MATRIX_SIZE
    """

    if not matrix:
        raise ValueError("Matrix cannot be null")
    
    if not isinstance(matrix, list):
        raise TypeError("Matrix must be a list")
    
    for i, row in enumerate(matrix):
        if not isinstance(row, list):
            raise ValueError(f"Row {i} must be a list")
    
    rows = len(matrix)
    for i in range (0, rows):
        if len(matrix[i]) != rows:
            raise ValueError("Matrix must be square")
        
    if rows > MAX_MATRIX_SIZE:
        raise ValueError(f"Matrix must have no more than {MAX_MATRIX_SIZE} columns")
        
    for i in range(rows):
        for j in range(rows):
            if not (matrix[i][j] == 0 or matrix[i][j] == 1):
                raise ValueError("Matrix can only contain 0s and 1s")

    return True

def print_matrix(matrix: List[List[int]]) -> None:
    """Nicely formats and prints a 2D matrix."""
    for row in matrix:
        print(" ".join(f"{cell:2}" for cell in row))

if __name__ == "__main__":
    not_a_list = minesweeper("TEST NOT A LIST") # Object must be a list

    not_all_rows_are_lists = minesweeper(
    [
        [1, 1, 1],
        [1, 1, 1],
        "THIS IS NOT A LIST"
    ]) # All the rows must be lists

    matrix_is_not_square = minesweeper(
    [
        [0, 1, 1],
        [1, 1, 1],
    ]) # Matrix must be square

    matrix_contains_unexpected_characters = minesweeper(
    [
        [0, 1, 1],
        [1, 1, 1],
        [1, '@', 1],
    ]) # Matriz can only contain 0 and 1s

    result = minesweeper(
    [
        [1, 1, 1, 1, 0],
        [1, 1, 1, 0, 1],
        [1, 0, 1, 1, 1],
        [0, 0, 1, 0, 0],
        [0, 1, 0, 1, 0],
    ])

    if result is not None:
        print_matrix(result)
