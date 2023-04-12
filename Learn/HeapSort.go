package main

import "fmt"

func main() {
	//Есть срез данных. Куча неотсортированная
	//Но мы хотим отсортировать её как бинарное дерево
	a := []int{7, 18, 3, 5, 42, 102, 17, 11, 0, 9}
	// Условно будем делить массив на 2 части:
	// Слева отсортированная куча, справа - остальные данные
	// Постепенно увеличиваем кучу и уменьшаем остаток
	n := len(a)
	var index int
	// Берём i-тый элемент и просеиваем вверх
	for i := 0; i < n; i++ {
		index = i
		for index-1 >= 0 && a[index] < a[(index-1)/2] {
			a[index], a[(index-1)/2] = a[(index-1)/2], a[index]
			index = (index - 1) / 2
		}
	}
	// Ура, получили кучу
	fmt.Println(a)

	// Теперь сортируем по убыванию кучу
	// Для этого берём 0-ой элемент, меняем его местами с крайним правым элементом
	// Новый 0-ой элемент просеиваем вниз
	// Теперь у нас снова есть куча слева и минимальный элемент справа
	// Снова меняем местами первый и последний элемент кучи
	//(величина кучи стала меньше после перестановки начального элемента)
	// И снова просеиваем вниз и формируем кучу
	for i := n - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		index = 0
		// Пока есть хотя бы 1 ребёнок и элемент меньше родителя, своп
		for index*2+1 < i {
			j := 2*index + 1
			//Если есть второй ребёнок и он меньше, чем первый, запоминаем его индекс как меньшего дитя
			if j+1 < i && a[j] > a[j+1] {
				j++
			}
			// Если наш элемент меньше меньшего ребёнка, то выходим, если нет - своп
			if a[index] <= a[j] {
				break
			} else {
				a[index], a[j] = a[j], a[index]
				index = j
			}
		}
	}
	// Ура, сортированный массив
	fmt.Println(a)

}
