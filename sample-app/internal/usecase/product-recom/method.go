package productrecom

// SetData set data for single user result
func (s *SingleUserResult) SetData(productRecom []ProductRecom) (affinity float32) {
	s.Data = productRecom
	s.SetOverallAffinities()
	return
}

// SetOverallAffinities set overall convidence for single user result
func (s *SingleUserResult) SetOverallAffinities() (affinity float32) {
	if s.Data == nil && len(s.Data) <= 0 {
		return
	}
	var total float32
	for _, productRecom := range s.Data {
		total += productRecom.Affinity
	}
	affinity = total / float32(len(s.Data))
	s.Meta.OverallAffinities = affinity
	return
}
