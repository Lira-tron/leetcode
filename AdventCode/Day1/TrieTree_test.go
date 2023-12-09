package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieTree_AddWord(t *testing.T) {
	type fields struct {
		r *Node
	}
	type args struct {
		w     string
		value int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected int
	}{
		{
			name: "Example 1",
			fields: fields{
				r: &Node{
					Neighbors:   map[byte]*Node{},
					Value:       0,
					IsEndOfWord: false,
				},
			},
			args: args{
				w:     "one",
				value: 1,
			},
			expected: 1,
		},
		{
			name: "Example 2",
			fields: fields{
				r: &Node{
					Neighbors:   map[byte]*Node{},
					Value:       0,
					IsEndOfWord: false,
				},
			},
			args: args{
				w:     "two",
				value: 2,
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TrieTree{
				r: tt.fields.r,
			}
			tr.AddWord(tt.args.w, tt.args.value)
			f, _ := tr.Find(tt.args.w)
			assert.Equal(t, f, tt.expected)
		})
	}
}

func TestTrieTree_RemoveWord(t *testing.T) {
	tr := &TrieTree{
		r: newNode(),
	}
	tr.AddWord("one", 1)
	tr.AddWord("two", 2)
	tr.AddWord("twot", 3)
	ok := tr.RemoveWord("one")
	assert.True(t, ok)
	ok = tr.RemoveWord("one")
	assert.False(t, ok)
	_, ok = tr.Find("two")
	assert.True(t, ok)
	_, ok = tr.Find("twot")
	assert.True(t, ok)
	ok = tr.RemoveWord("twot")
	_, ok = tr.Find("twot")
	assert.False(t, ok)
	_, ok = tr.Find("two")
	assert.True(t, ok)
}
