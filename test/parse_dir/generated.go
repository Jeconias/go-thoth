package any

import "github.com/lab259/go-thoth/validators"

// Thoth validate
func (h *Home) Thoth() (err error) {
	err = validators.IsInt(h.ID)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// Thoth validate
func (c *Client) Thoth() (err error) {
	err = validators.IsInt64(c.ID)
	if err != nil {
		return err
	}

	err = validators.Empty(c.Name)
	if err != nil {
		return err
	}

	err = validators.EmptyPtr(c.LastName)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// Thoth validate
func (u *User) Thoth() (err error) {
	err = validators.Empty(u.Name)
	if err != nil {
		return err
	}

	return nil
}

// Thoth validate
func (n *Number) Thoth() (err error) {
	err = validators.IsUint(n.A)
	if err != nil {
		return err
	}

	err = validators.IsUint8(n.B)
	if err != nil {
		return err
	}

	err = validators.IsUint16(n.C)
	if err != nil {
		return err
	}

	err = validators.IsUint32(n.D)
	if err != nil {
		return err
	}

	err = validators.IsUint64(n.E)
	if err != nil {
		return err
	}

	err = validators.IsUintptr(n.F)
	if err != nil {
		return err
	}

	err = validators.IsInt(n.G)
	if err != nil {
		return err
	}

	err = validators.IsInt8(n.H)
	if err != nil {
		return err
	}

	err = validators.IsInt16(n.I)
	if err != nil {
		return err
	}

	err = validators.IsInt32(n.J)
	if err != nil {
		return err
	}

	err = validators.IsInt64(n.K)
	if err != nil {
		return err
	}

	err = validators.IsFloat32(n.L)
	if err != nil {
		return err
	}

	err = validators.IsFloat64(n.M)
	if err != nil {
		return err
	}

	err = validators.IsComplex64(n.N)
	if err != nil {
		return err
	}

	err = validators.IsComplex128(n.O)
	if err != nil {
		return err
	}

	return nil
}

// Thoth validate
func (s *SliceString) Thoth() (err error) {
	err = validators.IsSliceString(s.A)
	if err != nil {
		return err
	}

	return nil
}
