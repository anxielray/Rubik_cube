package models


type Mid_cubit struct {
	Top_face string
	Front_face string
}

type Corner_cubit struct {
	Top_face string
	Front_face string
	Right_face string
}

type Top_Layer  struct {
	Left_front Corner_cubit
	Mid_front Mid_cubit
	Right_front Corner_cubit
	Right_mid Mid_cubit
	Right_back Corner_cubit
	Mid_back Mid_cubit
	Left_back Corner_cubit
	Left_mid Mid_cubit
	Center_cubit string
}

type Bottom_Layer  struct {
	Left_front Corner_cubit
	Mid_front Mid_cubit
	Right_front Corner_cubit
	Right_mid Mid_cubit
	Right_back Corner_cubit
	Mid_back Mid_cubit
	Left_back Corner_cubit
	Left_mid Mid_cubit
	Center_cubit string
}

type Middle_Layer  struct {
	Left_front Mid_cubit
	Center_front_cubit string
	Right_front Mid_cubit
	Center_right_cubit string
	Right_back Mid_cubit
	Center_back_cubit string
	Left_back Mid_cubit
	Center_left_cubit string
}

type Rubik_cube struct {
	Top_Layer Top_Layer
	Middle_Layer Middle_Layer
	Bottom_Layer Bottom_Layer
	Front_face string
}


var Commands []string