package models


func (r *Rubik_cube)Reshape_3D(color string) Rubik_cube {

}

func (r *Rubik_cube)U() Rubik_cube {

	tmp_r := *r

	//===( Update the frot row )===
	tmp_r.Top_Layer.Left_front = r.Top_Layer.Right_front
	tmp_r.Top_Layer.Mid_front = r.Top_Layer.Right_mid
	tmp_r.Top_Layer.Right_front = r.Top_Layer.Right_back

	//===( Update the right row )===
	tmp_r.Top_Layer.Right_mid = r.Top_Layer.Mid_back
	tmp_r.Top_Layer.Right_back = r.Top_Layer.Left_back

	//==( Update the back row )===
	tmp_r.Top_Layer.Mid_back = r.Top_Layer.Left_mid
	tmp_r.Top_Layer.Left_back = r.Top_Layer.Left_front

	//==( Update the left row )===
	tmp_r.Top_Layer.Left_mid = r.Top_Layer.Mid_front
	r = &tmp_r
	return *r
}

//===( This function will define the method that performs the operation of the rotation of the top layer in the counterclockwise direction )===
func (r *Rubik_cube)U_p() Rubik_cube {

	tmp_r := *r

	tmp_r.Top_Layer.Right_front = r.Top_Layer.Left_front
	tmp_r.Top_Layer.Right_mid = r.Top_Layer.Mid_front
	tmp_r.Top_Layer_Right_back = r.Top_Layer.Right_front

	tmp_r.Top_Layer.Mid_back = r.To_Layer.Right_mid
	tmp_r.Top_Layer.Left_back = r.Top_Layer.Right_back

	tmp_r.Top_Layer.Left_mid = r.Top_Layer.Mid_back
	tmp_r.Top_Layer.Left_front = r.Top_Layer.Left_back

	tmp_r.Top_Layer.Mid_front = r.Top_Layer.Left_mid
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)R() Rubik_cube {

	tmp_r := *r

	//==( Update the front column )===
	tm_r.Top_Layer.Right_front = r.Bottom_Layer.Right_front
	tmp_r.Middle_Layer.Right_front = r.Bottom_Layer.Right_mid
	tmp_r.Bottom_Layer.Right_front = r.Bottom_Layer_Right_back

	//==( Update the bottom column )===
	tmp_r.Bottom_Layer.Right_mid = r.Middle_Layer.Right_back
	tmp_r.Bottom_Layer.Right_back = r.Top_Layer.Right_back

	//==( Update the back column )===
	tmp_r.Middle_Layer.Right_back = r.Top_Layer.Right_mid
	tmp_r.Top_Layer.Right_back = r.Top_Layer.Right_front

	//==( Update the top column )===
	tmp_r.Top_Layer.Right_mid = r.Middle_Layer.Right_front
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)R_p() Rubik_cube {

	tmp_r := *r
	
	//==( Update the front column )===
	tm_r.Bottom_Layer.Right_front = r.Top_Layer.Right_front
	tmp_r.Bottom_Layer.Right_mid = r.Middle_Layer.Right_front
	tmp_r.Bottom_Layer_Right_back = r.Bottom_Layer.Right_front

	//==( Update the bottom column )===
	tmp_r.Middle_Layer.Right_back = r.Bottom_Layer.Right_mid
	tmp_r.Top_Layer.Right_back = r.Bottom_Layer.Right_back

	//==( Update the back column )===
	tmp_r.Top_Layer.Right_mid = r.Middle_Layer.Right_back
	tmp_r.Top_Layer.Right_front = r.Top_Layer.Right_back

	//==( Update the top column )===
	tmp_r.Middle_Layer.Right_front = r.Top_Layer.Right_mid
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)L() Rubik_cube {

	tmp_r := *r

	tmp_r.Top_Layer.Left_front = r.Top_Layer.Left_back
	tmp_r.Middle_Layer.Left_front = r.Top_Layer.Left_mid
	tmp_r.Bottom_Layer.Left_front = r.Top_Layer.Left_front

	tmp_r.Bottom_Layer.Left_mid = r.Middle_Layer.Left_front
	tmp_r.Bottom_Layer.Left_back = r.Bottom_Layer.Left_front

	tmp_r.Middle_Layer.Left_back = r.Bottom_Layer.Left_mid
	tmp_r.Top_Layer.Left_back = r.Bottom_Layer.Left_back

	tmp_r.Top_Layer.Left_mid = r.Middle_Layer.Left_back
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)L_p() Rubik_cube {

	tmp_r := *r

	tmp_r.Top_Layer.Left_back = r.Top_Layer.Left_front
	tmp_r.Top_Layer.Left_mid = r.Middle_Layer.Left_front
	tmp_r.Top_Layer.Left_front = r.Bottom_Layer.Left_front

	tmp_r.Middle_Layer.Left_front = r.Bottom_Layer.Left_mid
	tmp_r.Bottom_Layer.Left_front = r.Bottom_Layer.Left_back

	tmp_r.Bottom_Layer.Left_mid = r.Middle_Layer.Left_back
	tmp_r.Bottom_Layer.Left_back= r.Top_Layer.Left_back 

	tmp_r.Middle_Layer.Left_back = r.Top_Layer.Left_mid
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)F() Rubik_cube {

	tmp_r := *r

	tmp_r.Top_Layer.Left_front = r.Bottom_Layer_Left_front
	tmp_r.Top_Layer.Mid_front = r.Middle_Layer.Left_front
	tmp_r.Top_Layer.Right_front = r.Top_Layer.Left_front
	
	tmp_r.Middle_Layer.Left_front = r.Bottom_Layer.Mid_front
	tmp_r.Middle_Layer.Right_front = r.Top_Layer.Mid_front

	tmp_r.Bottom_Layer.Left_front = r.Bottom_Layer.Left_front
	tmp_r.Bottom_Layer.Mid_front = r.Middle_Layer.Right_front
	tmp_r.Bottom_Layer.Right_front = r.Top_Layer.Right_front
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)F_p() Rubik_cube {

	tmp_r := *r
	
	tmp_r.Bottom_Layer_Left_front = r.Top_Layer.Left_front
	tmp_r.Middle_Layer.Left_front = r.Top_Layer.Mid_front
	tmp_r.Top_Layer.Left_front = r.Top_Layer.Right_front
	
	tmp_r.Bottom_Layer.Mid_front = r.Middle_Layer.Left_front
	tmp_r.Top_Layer.Mid_front = r.Middle_Layer.Right_front

	tmp_r.Bottom_Layer.Left_front = r.Bottom_Layer.Left_front
	tmp_r.Middle_Layer.Right_front = r.Bottom_Layer.Mid_front
	tmp_r.Top_Layer.Right_front = r.Bottom_Layer.Right_front
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)B() Rubik_cube {

	tmp_r := *r

	tmp_r.Bottom_Layer_Left_back = r.Top_Layer.Left_back
	tmp_r.Middle_Layer.Left_back = r.Top_Layer.Mid_back
	tmp_r.Top_Layer.Left_back = r.Top_Layer.Right_back
	
	tmp_r.Bottom_Layer.Mid_back = r.Middle_Layer.Left_back
	tmp_r.Top_Layer.Mid_back = r.Middle_Layer.Right_back

	tmp_r.Bottom_Layer.Left_back = r.Bottom_Layer.Left_back
	tmp_r.Middle_Layer.Right_back = r.Bottom_Layer.Mid_back
	tmp_r.Top_Layer.Right_back = r.Bottom_Layer.Right_back
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)B_p() Rubik_cube {

	tmp_r := *r

	tmp_r.Top_Layer.Left_back = r.Bottom_Layer_Left_back
	tmp_r.Top_Layer.Mid_back = r.Middle_Layer.Left_back
	tmp_r.Top_Layer.Right_back = r.Top_Layer.Left_back
	
	tmp_r.Middle_Layer.Left_back = r.Bottom_Layer.Mid_back
	tmp_r.Middle_Layer.Right_back = r.Top_Layer.Mid_back

	tmp_r.Bottom_Layer.Left_back = r.Bottom_Layer.Left_back
	tmp_r.Bottom_Layer.Mid_back = r.Middle_Layer.Right_back
	tmp_r.Bottom_Layer.Right_back = r.Top_Layer.Right_back
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)D() Rubik_cube {

	tmp_r := *r

	tmp_r.Bottom_Layer.Right_front = r.Bottom_Layer.Left_front
	tmp_r.Bottom_Layer.Right_mid = r.Bottom_Layer.Mid_front
	tmp_r.Bottom_Layer_Right_back = r.Bottom_Layer.Right_front

	tmp_r.Bottom_Layer.Mid_back = r.Bottom_Layer.Right_mid
	tmp_r.Bottom_Layer.Left_back = r.Bottom_Layer.Right_back

	tmp_r.Bottom_Layer.Left_mid = r.Bottom_Layer.Mid_back
	tmp_r.Bottom_Layer.Left_front = r.Bottom_Layer.Left_back

	tmp_r.Bottom_Layer.Mid_front = r.Bottom_Layer.Left_mid
	r = &tmp_r
	return *r
}
func (r *Rubik_cube)D_p() Rubik_cube {

	tmp_r := *r

	//===( Update the frot row )===
	tmp_r.Bottom_Layer.Left_front = r.Bottom_Layer.Right_front
	tmp_r.Bottom_Layer.Mid_front = r.Bottom_Layer.Right_mid
	tmp_r.Bottom_Layer.Right_front = r.Bottom_Layer.Right_back

	//===( Update the right row )===
	tmp_r.Bottom_Layer.Right_mid = r.Bottom_Layer.Mid_back
	tmp_r.Bottom_Layer.Right_back = r.Bottom_Layer.Left_back

	//==( Update the back row )===
	tmp_r.Bottom_Layer.Mid_back = r.Bottom_Layer.Left_mid
	tmp_r.Bottom_Layer.Left_back = r.Bottom_Layer.Left_front

	//==( Update the left row )===
	tmp_r.BottomLayer.Left_mid = r.Bottom_Layer.Mid_front
	r = &tmp_r
	return *r
}
