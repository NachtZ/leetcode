
func rotate(matrix [][]int)  {
    size := len(matrix)
    if size <=1 {
        return
    }
    layers := size/2
    for layer:=0;layer<layers;layer++{
        for i:=layer;i<size -1 -layer;i++{
            temp := matrix[i][layer]
            matrix[i][layer] = matrix[size-1-layer][i]
            matrix[size-1-layer][i] = matrix[size-1-i][size-1-layer]
            matrix[size-1-i][size-1-layer] = matrix[layer][size-1-i]
            matrix[layer][size-1-i] = temp
        }
    }
}