from typing import List

"""
1. ¿La matriz es cuadrada? ¿o puede tener diferente número de filas y columnas? ✓ siempre es cuadrada
2. ¿El input qué tipo de dato es? ✓ ya se indica en el doc
3. ¿Qué debería pasar si en la matriz alguno de los datos no es cero ni uno? ✓ cero significa sin minas y cualquier cosa distinta de cero es una mina
4. ¿Qué debería pasar si no todas las filas tienen el mismo número de columnas? ✓ solucionado
_______________________________________________________________________________
Tipos de situaciones de error
* Número de columnas no es consistente entre todas las filas
_______________________________________________________________________________
* La matriz tiene un numero máximo de filas?

"""
MAX_MATRIX_SIZE = 50
# TODO no olvidar revisar si debería haber una enumeración en los numeros quemados

def minesweeper(matrix: List[List[int]]) -> List[List[int]]:
    try:
        if validate_matrix(matrix): 
            rows = len(matrix)
            matrix_response = [[0] * rows for _ in range(rows)]

            for i in range(rows):
                for j in range(rows):

                    if matrix[i][j] == 1:
                        matrix_response[i][j] = 9
                    else:
                        mine_sum = check_neighbors(i, j, rows, matrix)     
                        matrix_response[i][j] = mine_sum
            return matrix_response
    except ValueError as v:
        print(f"Value error: {v}")
    except TypeError as t:
        print(f"Type error: {t}")

def check_neighbors(i: int, j: int, rows: int, matrix: List[List[int]]) -> int:
    for k in range(-1, 2):
        for l in range(-1,2):
            if k == 0 and l == 0:
                continue
            if (i + k >= 0 and i + k <= rows -1) and (j + l >= 0 and j + l <= rows -1) and matrix[i+k][j+l] == 1:
                mine_sum = mine_sum + 1
    
def validate_matrix(matrix: List[List[int]]) -> bool:

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
        raise ValueError("Matrix must have no more than 50 columns")
        
    for i in range(rows):
        for j in range(rows):
            if not (matrix[i][j] == 0 or matrix[i][j] == 1):
                raise ValueError("Matrix can only contain 0s and 1s")

    return True

def print_matrix(matrix: List[List[int]]) -> None:
    for row in matrix:
        print(" ".join(f"{cell:2}" for cell in row))

if __name__ == "__main__":
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
