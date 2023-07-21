package type_constant

type SampleType string

func (e SampleType) String() string {
	return string(e)
}

const (
	SampleType_SAMPLE_1 SampleType = "SAMPLE_1"
	SampleType_SAMPLE_2 SampleType = "SAMPLE_2"
	SampleType_SAMPLE_3 SampleType = "SAMPLE_3"
)
