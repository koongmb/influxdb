// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: file_store.gen.go.tmpl

package tsm1

// ReadFloatBlock reads the next block as a set of float values.
func (c *KeyCursor) ReadFloatBlock(tdec *TimeDecoder, vdec *FloatDecoder, buf *[]FloatValue) ([]FloatValue, error) {
	// No matching blocks to decode
	if len(c.current) == 0 {
		return nil, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	*buf = (*buf)[:0]
	values, err := first.r.ReadFloatBlockAt(&first.entry, tdec, vdec, buf)

	// Remove values we already read
	values = FloatValues(values).Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	values = c.filterFloatValues(tombstones, values)

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if len(values) > 0 {
			first.markRead(values[0].UnixNano(), values[len(values)-1].UnixNano())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := values[0].UnixNano(), values[len(values)-1].UnixNano()
	if c.ascending {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				maxT = cur.entry.MaxTime
				values = FloatValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)
			var a []FloatValue
			v, err := cur.r.ReadFloatBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterFloatValues(tombstones, v)

			// Remove values we already read
			v = FloatValues(v).Exclude(cur.readMin, cur.readMax)

			if len(v) > 0 {
				// Only use values in the overlapping window
				v = FloatValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				// Merge the remaing values with the existing
				values = FloatValues(values).Merge(v)
			}
		}

	} else {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				minT = cur.entry.MinTime
				values = FloatValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)

			var a []FloatValue
			v, err := cur.r.ReadFloatBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterFloatValues(tombstones, v)

			// Remove values we already read
			v = FloatValues(v).Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if len(v) > 0 {
				v = FloatValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				values = FloatValues(v).Merge(values)
			}
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

// ReadIntegerBlock reads the next block as a set of integer values.
func (c *KeyCursor) ReadIntegerBlock(tdec *TimeDecoder, vdec *IntegerDecoder, buf *[]IntegerValue) ([]IntegerValue, error) {
	// No matching blocks to decode
	if len(c.current) == 0 {
		return nil, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	*buf = (*buf)[:0]
	values, err := first.r.ReadIntegerBlockAt(&first.entry, tdec, vdec, buf)

	// Remove values we already read
	values = IntegerValues(values).Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	values = c.filterIntegerValues(tombstones, values)

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if len(values) > 0 {
			first.markRead(values[0].UnixNano(), values[len(values)-1].UnixNano())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := values[0].UnixNano(), values[len(values)-1].UnixNano()
	if c.ascending {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				maxT = cur.entry.MaxTime
				values = IntegerValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)
			var a []IntegerValue
			v, err := cur.r.ReadIntegerBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterIntegerValues(tombstones, v)

			// Remove values we already read
			v = IntegerValues(v).Exclude(cur.readMin, cur.readMax)

			if len(v) > 0 {
				// Only use values in the overlapping window
				v = IntegerValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				// Merge the remaing values with the existing
				values = IntegerValues(values).Merge(v)
			}
		}

	} else {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				minT = cur.entry.MinTime
				values = IntegerValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)

			var a []IntegerValue
			v, err := cur.r.ReadIntegerBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterIntegerValues(tombstones, v)

			// Remove values we already read
			v = IntegerValues(v).Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if len(v) > 0 {
				v = IntegerValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				values = IntegerValues(v).Merge(values)
			}
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

// ReadStringBlock reads the next block as a set of string values.
func (c *KeyCursor) ReadStringBlock(tdec *TimeDecoder, vdec *StringDecoder, buf *[]StringValue) ([]StringValue, error) {
	// No matching blocks to decode
	if len(c.current) == 0 {
		return nil, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	*buf = (*buf)[:0]
	values, err := first.r.ReadStringBlockAt(&first.entry, tdec, vdec, buf)

	// Remove values we already read
	values = StringValues(values).Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	values = c.filterStringValues(tombstones, values)

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if len(values) > 0 {
			first.markRead(values[0].UnixNano(), values[len(values)-1].UnixNano())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := values[0].UnixNano(), values[len(values)-1].UnixNano()
	if c.ascending {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				maxT = cur.entry.MaxTime
				values = StringValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)
			var a []StringValue
			v, err := cur.r.ReadStringBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterStringValues(tombstones, v)

			// Remove values we already read
			v = StringValues(v).Exclude(cur.readMin, cur.readMax)

			if len(v) > 0 {
				// Only use values in the overlapping window
				v = StringValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				// Merge the remaing values with the existing
				values = StringValues(values).Merge(v)
			}
		}

	} else {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				minT = cur.entry.MinTime
				values = StringValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)

			var a []StringValue
			v, err := cur.r.ReadStringBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterStringValues(tombstones, v)

			// Remove values we already read
			v = StringValues(v).Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if len(v) > 0 {
				v = StringValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				values = StringValues(v).Merge(values)
			}
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

// ReadBooleanBlock reads the next block as a set of boolean values.
func (c *KeyCursor) ReadBooleanBlock(tdec *TimeDecoder, vdec *BooleanDecoder, buf *[]BooleanValue) ([]BooleanValue, error) {
	// No matching blocks to decode
	if len(c.current) == 0 {
		return nil, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	*buf = (*buf)[:0]
	values, err := first.r.ReadBooleanBlockAt(&first.entry, tdec, vdec, buf)

	// Remove values we already read
	values = BooleanValues(values).Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	values = c.filterBooleanValues(tombstones, values)

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if len(values) > 0 {
			first.markRead(values[0].UnixNano(), values[len(values)-1].UnixNano())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := values[0].UnixNano(), values[len(values)-1].UnixNano()
	if c.ascending {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				maxT = cur.entry.MaxTime
				values = BooleanValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)
			var a []BooleanValue
			v, err := cur.r.ReadBooleanBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterBooleanValues(tombstones, v)

			// Remove values we already read
			v = BooleanValues(v).Exclude(cur.readMin, cur.readMax)

			if len(v) > 0 {
				// Only use values in the overlapping window
				v = BooleanValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				// Merge the remaing values with the existing
				values = BooleanValues(values).Merge(v)
			}
		}

	} else {
		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				minT = cur.entry.MinTime
				values = BooleanValues(values).Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				continue
			}

			tombstones := cur.r.TombstoneRange(c.key)

			var a []BooleanValue
			v, err := cur.r.ReadBooleanBlockAt(&cur.entry, tdec, vdec, &a)
			if err != nil {
				return nil, err
			}
			// Remove any tombstoned values
			v = c.filterBooleanValues(tombstones, v)

			// Remove values we already read
			v = BooleanValues(v).Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if len(v) > 0 {
				v = BooleanValues(v).Include(minT, maxT)

				if len(v) > 0 {
					cur.markRead(v[0].UnixNano(), v[len(v)-1].UnixNano())
				}
				values = BooleanValues(v).Merge(values)
			}
		}
	}

	first.markRead(minT, maxT)

	return values, err
}
