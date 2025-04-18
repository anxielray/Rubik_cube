package models

func (r *Rubik_cube) U() *Rubik_cube {

	tmp_r := *r
	tmp_r.Top_Layer = Top_Layer{
		Left_front:  Corner_cubit{Top_face: r.Top_Layer.Right_front.Top_face, Front_face: r.Top_Layer.Right_front.Side_face, Side_face: r.Top_Layer.Right_front.Front_face},
		Left_mid:    Mid_cubit{Top_face: r.Top_Layer.Mid_front.Top_face, Front_face: r.Top_Layer.Mid_front.Front_face},
		Left_back:   Corner_cubit{Top_face: r.Top_Layer.Left_front.Top_face, Front_face: r.Top_Layer.Left_front.Side_face, Side_face: r.Top_Layer.Left_front.Front_face},
		Mid_front:   Mid_cubit{Top_face: r.Top_Layer.Right_mid.Top_face, Front_face: r.Top_Layer.Right_mid.Front_face},
		Mid_back:    Mid_cubit{Top_face: r.Top_Layer.Left_mid.Top_face, Front_face: r.Top_Layer.Left_mid.Front_face},
		Right_front: Corner_cubit{Top_face: r.Top_Layer.Right_back.Top_face, Front_face: r.Top_Layer.Right_back.Front_face, Side_face: r.Top_Layer.Right_back.Side_face},
		Right_mid:   Mid_cubit{Top_face: r.Top_Layer.Mid_back.Top_face, Front_face: r.Top_Layer.Mid_back.Front_face},
		Right_back:  Corner_cubit{Top_face: r.Top_Layer.Left_back.Top_face, Front_face: r.Top_Layer.Left_back.Front_face, Side_face: r.Top_Layer.Left_back.Side_face},
	}

	r = &tmp_r
	return r
}

// ===( This function will define the method that performs the operation of the rotation of the top layer in the counterclockwise direction )===
func (r *Rubik_cube) U_p() *Rubik_cube {

	tmp_r := *r
	tmp_r.Top_Layer = Top_Layer{
		Left_front:  Corner_cubit{Top_face: r.Top_Layer.Left_back.Top_face, Front_face: r.Top_Layer.Left_back.Side_face, Side_face: r.Top_Layer.Left_back.Front_face},
		Left_mid:    Mid_cubit{Top_face: r.Top_Layer.Mid_back.Top_face, Front_face: r.Top_Layer.Mid_front.Front_face},
		Left_back:   Corner_cubit{Top_face: r.Top_Layer.Right_back.Top_face, Front_face: r.Top_Layer.Right_back.Front_face, Side_face: r.Top_Layer.Right_back.Side_face},
		Mid_front:   Mid_cubit{Top_face: r.Top_Layer.Left_mid.Top_face, Front_face: r.Top_Layer.Left_mid.Front_face},
		Mid_back:    Mid_cubit{Top_face: r.Top_Layer.Right_mid.Top_face, Front_face: r.Top_Layer.Right_mid.Front_face},
		Right_front: Corner_cubit{Top_face: r.Top_Layer.Left_front.Top_face, Front_face: r.Top_Layer.Left_front.Side_face, Side_face: r.Top_Layer.Left_front.Front_face},
		Right_mid:   Mid_cubit{Top_face: r.Top_Layer.Mid_front.Top_face, Front_face: r.Top_Layer.Mid_front.Front_face},
		Right_back:  Corner_cubit{Top_face: r.Top_Layer.Right_front.Top_face, Front_face: r.Top_Layer.Right_front.Front_face, Side_face: r.Top_Layer.Right_front.Side_face},
	}
	r = &tmp_r
	return r
}

func (r *Rubik_cube) R() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  r.Top_Layer.Left_front,
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   r.Top_Layer.Left_back,
			Mid_front:   r.Top_Layer.Mid_front,
			Mid_back:    r.Top_Layer.Mid_back,
			Right_front: Corner_cubit{Top_face: r.Bottom_Layer.Right_front.Front_face, Front_face: r.Bottom_Layer.Right_front.Top_face, Side_face: r.Bottom_Layer.Right_front.Side_face},
			Right_mid:   Mid_cubit{Top_face: r.Middle_Layer.Right_front.Front_face, Front_face: r.Middle_Layer.Right_front.Top_face},
			Right_back:  Corner_cubit{Top_face: r.Top_Layer.Right_front.Front_face, Front_face: r.Top_Layer.Right_front.Side_face, Side_face: r.Top_Layer.Right_front.Top_face},
		},
		Middle_Layer: Middle_Layer{
			Left_front:         r.Middle_Layer.Left_front,
			Left_back:          r.Middle_Layer.Left_back,
			Right_front:        Mid_cubit{Front_face: r.Bottom_Layer.Right_mid.Top_face, Top_face: r.Bottom_Layer.Right_mid.Front_face},
			Right_back:         Mid_cubit{Front_face: r.Top_Layer.Right_mid.Front_face, Top_face: r.Top_Layer.Right_mid.Top_face},
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   r.Bottom_Layer.Left_front,
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    r.Bottom_Layer.Left_back,
			Mid_front:    r.Bottom_Layer.Mid_front,
			Mid_back:     r.Bottom_Layer.Mid_back,
			Right_front:  Corner_cubit{Top_face: r.Bottom_Layer.Right_back.Side_face, Front_face: r.Bottom_Layer.Right_back.Side_face, Side_face: r.Bottom_Layer.Right_back.Front_face},
			Right_mid:    Mid_cubit{Top_face: r.Middle_Layer.Right_back.Top_face, Front_face: r.Middle_Layer.Right_back.Front_face},
			Right_back:   Corner_cubit{Top_face: r.Top_Layer.Right_back.Side_face, Front_face: r.Top_Layer.Right_back.Front_face, Side_face: r.Top_Layer.Right_back.Top_face},
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) R_p() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  r.Top_Layer.Left_front,
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   r.Top_Layer.Left_back,
			Mid_front:   r.Top_Layer.Mid_front,
			Mid_back:    r.Top_Layer.Mid_back,
			Right_front: Corner_cubit{Top_face:  r.Top_Layer.Right_back.Side_face, Front_face: r.Top_Layer.Right_back.Top_face , Side_face: r.Top_Layer.Right_back.Front_face},
			Right_mid:   Mid_cubit{Top_face: r.Middle_Layer.Right_back.Top_face, Front_face: r.Middle_Layer.Right_back.Front_face},
			Right_back:  Corner_cubit{Top_face: r.Bottom_Layer.Right_back.Side_face, Front_face: r.Bottom_Layer.Right_back.Front_face, Side_face: r.Bottom_Layer.Right_back.Top_face},
		},
		Middle_Layer: Middle_Layer{
			Left_front:         r.Middle_Layer.Left_front,
			Left_back:          r.Middle_Layer.Left_back,
			Right_front:        Mid_cubit{Top_face: r.Top_Layer.Right_mid.Front_face, Front_face: r.Top_Layer.Right_mid.Top_face},
			Right_back:         Mid_cubit{Top_face: r.Bottom_Layer.Right_mid.Top_face, Front_face: r.Bottom_Layer.Right_mid.Front_face},
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   r.Bottom_Layer.Left_front,
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    r.Bottom_Layer.Left_back,
			Mid_front:    r.Bottom_Layer.Mid_front,
			Mid_back:     r.Bottom_Layer.Mid_back,
			Right_front:  Corner_cubit{Top_face: r.Top_Layer.Right_front.Front_face, Front_face: r.Top_Layer.Right_front.Top_face, Side_face: r.Top_Layer.Right_front.Side_face},
			Right_mid:    Mid_cubit{Top_face: r.Middle_Layer.Right_front.Front_face, Front_face: r.Middle_Layer.Right_front.Top_face},
			Right_back:   Corner_cubit{Top_face: r.Bottom_Layer.Right_front.Front_face, Front_face: r.Bottom_Layer.Right_front.Side_face, Side_face: r.Bottom_Layer.Right_front.Top_face},
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) L() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  Corner_cubit{Top_face: r.Top_Layer.Left_back.Front_face, Front_face: r.Top_Layer.Left_back.Top_face, Side_face: r.Top_Layer.Left_back.Side_face},
			Left_mid:    Mid_cubit{Front_face: r.Middle_Layer.Left_back.Top_face, Top_face: r.Middle_Layer.Left_back.Front_face},
			Left_back:   Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Front_face, Front_face: r.Bottom_Layer.Left_back.Top_face, Side_face: r.Bottom_Layer.Left_back.Side_face},
			Mid_front:   r.Top_Layer.Mid_front,
			Mid_back:    r.Top_Layer.Mid_back,
			Right_front: r.Top_Layer.Right_front,
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  r.Top_Layer.Right_back,
		},
		Middle_Layer: Middle_Layer{
			Left_front:         Mid_cubit{Top_face: r.Top_Layer.Left_mid.Front_face, Front_face: r.Top_Layer.Left_mid.Top_face},
			Left_back:          Mid_cubit{Top_face: r.Bottom_Layer.Left_mid.Front_face, Front_face: r.Bottom_Layer.Left_mid.Top_face},
			Right_front:        r.Middle_Layer.Right_front,
			Right_back:         r.Middle_Layer.Right_back,
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   Corner_cubit{Top_face: r.Top_Layer.Left_front.Front_face, Front_face: r.Top_Layer.Left_front.Top_face, Side_face: r.Top_Layer.Left_front.Side_face},
			Left_mid:     Mid_cubit{Top_face: r.Middle_Layer.Left_front.Front_face, Front_face: r.Middle_Layer.Left_front.Top_face},
			Left_back:    Corner_cubit{Top_face: r.Bottom_Layer.Left_front.Front_face, Front_face: r.Bottom_Layer.Left_front.Top_face, Side_face: r.Bottom_Layer.Left_front.Side_face},
			Mid_front:    r.Bottom_Layer.Mid_front,
			Mid_back:     r.Bottom_Layer.Mid_back,
			Right_front:  r.Bottom_Layer.Right_front,
			Right_mid:    r.Bottom_Layer.Right_mid,
			Right_back:   r.Bottom_Layer.Right_back,
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) L_p() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  Corner_cubit{Top_face: r.Bottom_Layer.Left_front.Front_face, Front_face: r.Bottom_Layer.Left_front.Top_face, Side_face: r.Bottom_Layer.Left_front.Side_face},
			Left_mid:    Mid_cubit{Top_face: r.Middle_Layer.Left_front.Front_face, Front_face: r.Middle_Layer.Left_front.Top_face},
			Left_back:   Corner_cubit{Top_face: r.Top_Layer.Left_front.Front_face, Front_face: r.Top_Layer.Left_front.Top_face , Side_face: r.Top_Layer.Left_front.Side_face},
			Mid_front:   r.Top_Layer.Mid_front,
			Mid_back:    r.Top_Layer.Mid_back,
			Right_front: r.Top_Layer.Right_front,
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  r.Top_Layer.Right_back,
		},
		Middle_Layer: Middle_Layer{
			Left_front:         Mid_cubit{Top_face: r.Bottom_Layer.Left_mid.Front_face, Front_face: r.Bottom_Layer.Left_mid.Top_face},
			Left_back:          Mid_cubit{Top_face: r.Top_Layer.Left_mid.Front_face, Front_face: r.Top_Layer.Left_mid.Top_face},
			Right_front:        r.Middle_Layer.Right_front,
			Right_back:         r.Middle_Layer.Right_back,
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Front_face, Front_face: r.Bottom_Layer.Left_back. Top_face, Side_face: r.Bottom_Layer.Left_back.Side_face},
			Left_mid:     Mid_cubit{Top_face: r.Middle_Layer.Left_back.Front_face, Front_face: r.Middle_Layer.Left_back.Top_face},
			Left_back:    Corner_cubit{Top_face: r.Top_Layer.Left_back.Front_face, Front_face: r.Top_Layer.Left_back.Top_face, Side_face: r.Top_Layer.Left_back.Side_face},
			Mid_front:    r.Bottom_Layer.Mid_front,
			Mid_back:     r.Bottom_Layer.Mid_back,
			Right_front:  r.Bottom_Layer.Right_front,
			Right_mid:    r.Bottom_Layer.Right_mid,
			Right_back:   r.Bottom_Layer.Right_back,
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) F() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  Corner_cubit{Top_face: r.Bottom_Layer.Left_front.Side_face, Front_face: r.Bottom_Layer.Left_front.Front_face, Side_face: r.Bottom_Layer.Left_front.Top_face},
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   r.Top_Layer.Left_back,
			Mid_front:   Mid_cubit{Top_face: r.Middle_Layer.Left_front.Top_face, Front_face: r.Middle_Layer.Left_front.Front_face},
			Mid_back:    r.Top_Layer.Mid_back,
			Right_front: Corner_cubit{Top_face: r.Top_Layer.Left_front.Side_face, Front_face: r.Top_Layer.Left_front.Front_face, Side_face: r.Top_Layer.Left_front.Top_face},
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  r.Top_Layer.Right_back,
		},
		Middle_Layer: Middle_Layer{
			Left_front:         Mid_cubit{Top_face: r.Bottom_Layer.Mid_front.Top_face, Front_face: r.Bottom_Layer.Mid_front.Front_face},
			Left_back:          r.Middle_Layer.Left_back,
			Right_front:        Mid_cubit{Top_face: r.Top_Layer.Mid_front.Top_face, Front_face: r.Top_Layer.Mid_front.Front_face},
			Right_back:         r.Middle_Layer.Right_back,
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   Corner_cubit{Top_face: r.Bottom_Layer.Right_front.Side_face, Front_face: r.Bottom_Layer.Right_front.Front_face, Side_face: r.Bottom_Layer.Right_front.Top_face},
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    r.Bottom_Layer.Left_back,
			Mid_front:    Mid_cubit{Top_face: r.Middle_Layer.Right_front.Top_face, Front_face: r.Middle_Layer.Right_front.Front_face},
			Mid_back:     r.Bottom_Layer.Mid_back,
			Right_front:  Corner_cubit{Top_face: r.Top_Layer.Right_front.Side_face, Front_face: r.Top_Layer.Right_front.Front_face, Side_face: r.Top_Layer.Right_front.Top_face},
			Right_mid:    Mid_cubit{Top_face: r.Middle_Layer.Right_back.Top_face, Front_face: r.Middle_Layer.Right_back.Front_face},
			Right_back:   Corner_cubit{Top_face: r.Top_Layer.Right_back.Side_face, Front_face: r.Top_Layer.Right_back.Front_face, Side_face: r.Top_Layer.Right_back.Top_face},
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) F_p() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  Corner_cubit{Top_face: r.Top_Layer.Right_front.Side_face, Front_face: r.Top_Layer.Right_front.Front_face, Side_face: r.Top_Layer.Right_front.Top_face},
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   r.Top_Layer.Left_back,
			Mid_front:   Mid_cubit{Top_face: r.Middle_Layer.Right_front.Top_face, Front_face: r.Middle_Layer.Right_front.Front_face},
			Mid_back:    r.Top_Layer.Mid_back,
			Right_front: Corner_cubit{Top_face: r.Bottom_Layer.Right_front.Side_face, Front_face: r.Bottom_Layer.Right_front.Front_face, Side_face: r.Bottom_Layer.Right_front.Top_face},
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  r.Top_Layer.Right_back,
		},
		Middle_Layer: Middle_Layer{
			Left_front:         r.Top_Layer.Mid_front,
			Left_back:          r.Middle_Layer.Left_back,
			Right_front:        r.Bottom_Layer.Mid_front,
			Right_back:         r.Middle_Layer.Right_back,
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   Corner_cubit{Top_face: r.Top_Layer.Left_front.Side_face, Front_face: r.Top_Layer.Left_front.Front_face, Side_face: r.Top_Layer.Left_front.Top_face},
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    r.Bottom_Layer.Left_back,
			Mid_front:    r.Middle_Layer.Left_front,
			Mid_back:     r.Bottom_Layer.Mid_back,
			Right_front:  Corner_cubit{Top_face: r.Bottom_Layer.Left_front.Side_face, Front_face: r.Bottom_Layer.Left_front.Front_face, Side_face: r.Bottom_Layer.Left_front.Top_face},
			Right_mid:    r.Bottom_Layer.Right_mid,
			Right_back:   r.Bottom_Layer.Right_back,
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) B() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  r.Top_Layer.Left_front,
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   Corner_cubit{Top_face: r.Top_Layer.Right_back.Front_face, Front_face: r.Top_Layer.Right_back.Side_face , Side_face: r.Top_Layer.Right_back.Top_face},
			Mid_front:   r.Top_Layer.Mid_front,
			Mid_back:    r.Middle_Layer.Right_back,
			Right_front: r.Top_Layer.Right_front,
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  Corner_cubit{Top_face: r.Bottom_Layer.Right_back.Front_face, Front_face: r.Bottom_Layer.Right_back.Top_face, Side_face: r.Bottom_Layer.Right_back.Side_face},
		},
		Middle_Layer: Middle_Layer{
			Left_front:         r.Middle_Layer.Left_front,
			Left_back:          r.Top_Layer.Mid_back,
			Right_front:        r.Middle_Layer.Right_front,
			Right_back:         r.Bottom_Layer.Mid_back,
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   r.Bottom_Layer.Left_front,
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    Corner_cubit{Top_face: r.Top_Layer.Left_front.Side_face, Front_face: r.Top_Layer.Left_front.Front_face, Side_face: r.Top_Layer.Left_front.Top_face},
			Mid_front:    r.Bottom_Layer.Mid_front,
			Mid_back:     r.Middle_Layer.Left_back,
			Right_front: r.Bottom_Layer.Right_front,
			Right_mid:    r.Bottom_Layer.Right_mid,
			Right_back:   Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Side_face, Front_face: r.Bottom_Layer.Left_back.Top_face, Side_face: r.Bottom_Layer.Left_back.Front_face},
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) B_p() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  r.Top_Layer.Left_front,
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Side_face, Front_face: r.Bottom_Layer.Left_back.Front_face, Side_face: r.Bottom_Layer.Left_back.Top_face},
			Mid_front:   r.Top_Layer.Mid_front,
			Mid_back:    r.Middle_Layer.Left_back,
			Right_front: r.Top_Layer.Right_front,
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  Corner_cubit{Top_face: r.Top_Layer.Left_back.Side_face, Front_face: r.Top_Layer.Left_back.Front_face, Side_face: r.Top_Layer.Left_back.Top_face},
		},
		Middle_Layer: Middle_Layer{
			Left_front:         r.Middle_Layer.Left_front,
			Left_back:          r.Bottom_Layer.Mid_back,
			Right_front:        r.Middle_Layer.Right_front,
			Right_back:         r.Top_Layer.Mid_back,
			Center_front_cubit: r.Middle_Layer.Center_front_cubit,
			Center_back_cubit:  r.Middle_Layer.Center_back_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   r.Bottom_Layer.Left_front,
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Side_face, Front_face: r.Bottom_Layer.Left_back.Front_face, Side_face: r.Bottom_Layer.Left_back.Top_face},
			Mid_front:    r.Bottom_Layer.Mid_front,
			Mid_back:     r.Middle_Layer.Right_back,
			Right_front: r.Bottom_Layer.Right_front,
			Right_mid:    r.Bottom_Layer.Right_mid,
			Right_back:   Corner_cubit{Top_face: r.Top_Layer.Right_back.Front_face, Front_face: r.Top_Layer.Right_back.Top_face, Side_face: r.Top_Layer.Right_back.Side_face},
			Center_cubit: r.Bottom_Layer.Center_cubit,
		},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) D() *Rubik_cube {

	tmp_r := *r
	tmp_r.Bottom_Layer = Bottom_Layer{
		Left_front:  Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Top_face, Front_face: r.Bottom_Layer.Left_back.Side_face, Side_face: r.Bottom_Layer.Left_back.Front_face},
		Left_mid:    Mid_cubit{Top_face: r.Bottom_Layer.Mid_back.Top_face, Front_face: r.Bottom_Layer.Mid_back.Front_face},
		Left_back:   Corner_cubit{Top_face: r.Bottom_Layer.Right_back.Top_face, Front_face: r.Bottom_Layer.Right_back.Front_face, Side_face: r.Bottom_Layer.Right_back.Side_face},
		Mid_front:   Mid_cubit{Top_face: r.Bottom_Layer.Left_mid.Top_face, Front_face: r.Bottom_Layer.Left_mid.Front_face},
		Mid_back:    Mid_cubit{Top_face: r.Bottom_Layer.Right_mid.Top_face, Front_face: r.Bottom_Layer.Right_mid.Front_face},
		Right_front: Corner_cubit{Top_face: r.Bottom_Layer.Left_front.Top_face, Front_face: r.Bottom_Layer.Left_front.Front_face, Side_face: r.Bottom_Layer.Left_front.Side_face},
		Right_mid:   Mid_cubit{Top_face: r.Bottom_Layer.Mid_front.Top_face, Front_face: r.Bottom_Layer.Mid_front.Front_face},
		Right_back:  Corner_cubit{Top_face: r.Bottom_Layer.Right_front.Top_face, Front_face: r.Bottom_Layer.Right_front.Front_face, Side_face: r.Bottom_Layer.Right_front.Side_face},
	}
	r = &tmp_r
	return r
}
func (r *Rubik_cube) D_p() *Rubik_cube {

	tmp_r := *r
	tmp_r.Bottom_Layer = Bottom_Layer{
		Left_front:  Corner_cubit{Top_face: r.Bottom_Layer.Right_front.Top_face, Front_face: r.Bottom_Layer.Right_front.Side_face, Side_face: r.Bottom_Layer.Right_front.Front_face},
		Left_mid:    Mid_cubit{Top_face: r.Bottom_Layer.Mid_front.Top_face, Front_face: r.Bottom_Layer.Mid_front.Front_face},
		Left_back:   Corner_cubit{Top_face: r.Bottom_Layer.Left_front.Top_face, Front_face: r.Bottom_Layer.Left_front.Side_face, Side_face: r.Bottom_Layer.Left_front.Front_face},
		Mid_front:   Mid_cubit{Top_face: r.Bottom_Layer.Right_mid.Top_face, Front_face: r.Bottom_Layer.Right_mid.Front_face},
		Mid_back:    Mid_cubit{Top_face: r.Bottom_Layer.Left_mid.Top_face, Front_face: r.Bottom_Layer.Left_mid.Front_face},
		Right_front: Corner_cubit{Top_face: r.Bottom_Layer.Right_back.Top_face, Front_face: r.Bottom_Layer.Right_back.Front_face, Side_face: r.Bottom_Layer.Right_back.Side_face},
		Right_mid:   Mid_cubit{Top_face: r.Bottom_Layer.Mid_back.Top_face, Front_face: r.Bottom_Layer.Mid_back.Front_face},
		Right_back:  Corner_cubit{Top_face: r.Bottom_Layer.Left_back.Top_face, Front_face: r.Bottom_Layer.Left_back.Front_face, Side_face: r.Bottom_Layer.Left_back.Side_face},
	}
	r = &tmp_r
	return r
}

func (r *Rubik_cube) Mx() *Rubik_cube {

	tmp_r := *r
	tmp_r.Middle_Layer = Middle_Layer{
		Left_front: Mid_cubit{Top_face: r.Middle_Layer.Right_front.Front_face, Front_face: r.Middle_Layer.Right_front.Top_face},
		Center_left_cubit: r.Middle_Layer.Center_front_cubit,
		Left_back: Mid_cubit{Top_face: r.Middle_Layer.Left_front.Front_face, Front_face: r.Middle_Layer.Left_front.Top_face},
		Center_front_cubit: r.Middle_Layer.Center_right_cubit,
		Center_back_cubit: r.Middle_Layer.Center_left_cubit,
		Right_front: Mid_cubit{Top_face: r.Middle_Layer.Right_back.Top_face, Front_face: r.Middle_Layer.Right_back.Front_face},
		Center_right_cubit: r.Middle_Layer.Center_back_cubit,
		Right_back: Mid_cubit{Top_face: r.Middle_Layer.Left_back.Top_face, Front_face: r.Middle_Layer.Left_back.Front_face},
	}
	r = &tmp_r
	return r
}

func (r *Rubik_cube) My() *Rubik_cube {

	tmp_r := Rubik_cube{
		Top_Layer: Top_Layer{
			Left_front:  r.Top_Layer.Left_front,
			Left_mid:    r.Top_Layer.Left_mid,
			Left_back:   r.Top_Layer.Left_back,
			Mid_front:   Mid_cubit{Top_face: r.Bottom_Layer.Mid_front.Front_face, Front_face: r.Bottom_Layer.Mid_front.Top_face},
			Mid_back:    Mid_cubit{Top_face: r.Top_Layer.Mid_front.Front_face, Front_face: r.Top_Layer.Mid_front.Top_face},
			Right_front: r.Top_Layer.Right_front,
			Right_mid:   r.Top_Layer.Right_mid,
			Right_back:  r.Top_Layer.Right_back,
		},
		Middle_Layer: Middle_Layer{
			Left_front:         r.Middle_Layer.Left_front,
			Left_back:          r.Middle_Layer.Left_back,
			Right_front:        r.Middle_Layer.Right_front,
			Right_back:         r.Middle_Layer.Right_back,
			Center_front_cubit: r.Bottom_Layer.Center_cubit,
			Center_back_cubit:  r.Top_Layer.Center_cubit,
			Center_right_cubit: r.Middle_Layer.Center_right_cubit,
			Center_left_cubit:  r.Middle_Layer.Center_left_cubit,
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   r.Bottom_Layer.Left_front,
			Left_mid:     r.Bottom_Layer.Left_mid,
			Left_back:    r.Bottom_Layer.Left_back,
			Mid_front:    Mid_cubit{Top_face: r.Bottom_Layer.Mid_back.Front_face, Front_face: r.Bottom_Layer.Mid_back.Top_face},
			Mid_back:     Mid_cubit{Top_face: r.Top_Layer.Mid_back.Front_face, Front_face: r.Top_Layer.Mid_back.Top_face},
			Right_front:  r.Bottom_Layer.Right_front,
			Right_mid:    r.Bottom_Layer.Right_mid,
			Right_back:   r.Bottom_Layer.Right_back,
			Center_cubit: r.Middle_Layer.Center_back_cubit,
		},
	}
	r = &tmp_r
	return r
}
