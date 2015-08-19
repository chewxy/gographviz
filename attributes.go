//Copyright 2013 Vastech SA (PTY) LTD
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package gographviz

import "github.com/awalterschulze/gographviz/common"

// This exposes the common.Attribute attributes. So instead of having to import github.com/awalterschulze/gographviz/ and
// github.com/awalterschulze/gographviz/common, you can just import github.com/awalterschulze/gographviz/, and pass in these as
// the attributes
type Attribute common.Attribute

const (
	DAMPING Attribute = iota
	K
	URL
	_BACKGROUND
	AREA
	ARROWHEAD
	ARROWSIZE
	ARROWTAIL
	BB
	BGCOLOR
	CENTER
	CHARSET
	CLUSTERRANK
	COLOR
	COLORSCHEME
	COMMENT
	COMPOUND
	CONCENTRATE
	CONSTRAINT
	DECORATE
	DEFAULTDIST
	DIM
	DIMEN
	DIR
	DIREDGECONSTRAINTS
	DISTORTION
	DPI
	EDGEURL
	EDGEHREF
	EDGETARGET
	EDGETOOLTIP
	EPSILON
	ESEP
	FILLCOLOR
	FIXEDSIZE
	FONTCOLOR
	FONTNAME
	FONTNAMES
	FONTPATH
	FONTSIZE
	FORCELABELS
	GRADIENTANGLE
	GROUP
	HEADURL
	HEAD_LP
	HEADCLIP
	HEADHREF
	HEADLABEL
	HEADPORT
	HEADTARGET
	HEADTOOLTIP
	HEIGHT
	HREF
	ID
	IMAGE
	IMAGEPATH
	IMAGESCALE
	INPUTSCALE
	LABEL
	LABELURL
	LABEL_SCHEME
	LABELANGLE
	LABELDISTANCE
	LABELFLOAT
	LABELFONTCOLOR
	LABELFONTNAME
	LABELFONTSIZE
	LABELHREF
	LABELJUST
	LABELLOC
	LABELTARGET
	LABELTOOLTIP
	LANDSCAPE
	LAYER
	LAYERLISTSEP
	LAYERS
	LAYERSELECT
	LAYERSEP
	LAYOUT
	LEN
	LEVELS
	LEVELSGAP
	LHEAD
	LHEIGHT
	LP
	LTAIL
	LWIDTH
	MARGIN
	MAXITER
	MCLIMIT
	MINDIST
	MINLEN
	MODE
	MODEL
	MOSEK
	NODESEP
	NOJUSTIFY
	NORMALIZE
	NOTRANSLATE
	NSLIMIT
	NSLIMIT1
	ORDERING
	ORIENTATION
	OUTPUTORDER
	OVERLAP
	OVERLAP_SCALING
	OVERLAP_SHRINK
	PACK
	PACKMODE
	PAD
	PAGE
	PAGEDIR
	PENCOLOR
	PENWIDTH
	PERIPHERIES
	PIN
	POS
	QUADTREE
	QUANTUM
	RANK
	RANKDIR
	RANKSEP
	RATIO
	RECTS
	REGULAR
	REMINCROSS
	REPULSIVEFORCE
	RESOLUTION
	ROOT
	ROTATE
	ROTATION
	SAMEHEAD
	SAMETAIL
	SAMPLEPOINTS
	SCALE
	SEARCHSIZE
	SEP
	SHAPE
	SHAPEFILE
	SHOWBOXES
	SIDES
	SIZE
	SKEW
	SMOOTHING
	SORTV
	SPLINES
	START
	STYLE
	STYLESHEET
	TAILURL
	TAIL_LP
	TAILCLIP
	TAILHREF
	TAILLABEL
	TAILPORT
	TAILTARGET
	TAILTOOLTIP
	TARGET
	TOOLTIP
	TRUECOLOR
	VERTICES
	VIEWPORT
	VORO_MARGIN
	WEIGHT
	WIDTH
	XDOTVERSION
	XLABEL
	XLP
	Z

	MAXATTRIBUTE
)

func fromStringMap(input map[string]string) (Attrs, bool) {
	attrs := NewAttrs()
	for k, v := range input {
		attr, ok := common.StringToAttribute(k)
		if !ok {
			return attrs, ok
		}
		attrs[attr] = v
	}
	return attrs, true
}