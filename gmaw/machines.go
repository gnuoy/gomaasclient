package gmaw

import (
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/maas"

	"github.com/juju/gomaasapi"
)

// Machines implements the maas.MachinesFetcher interface.
type Machines struct {
	client *gomaasapi.MAASObject
}

// NewMachines returns a pointer to a Machines.
func NewMachines(client *gomaasapi.MAASObject) *Machines {
	return &Machines{client: client}
}

// callPost returns the raw response from the MAAS API and any errors.
// This method creates the appropriate MAASObject for the API call, invokes the
// CallPost function, and returns the GetBytes() method of the response. It will
// return a nil byte array if CallPost returns an error.
func (m *Machines) callPost(op string, qsp url.Values) ([]byte, error) {
	mc := m.client.GetSubObject("machines")
	res, err := mc.CallPost(op, qsp)
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}

// Allocate fulfills the maas.MachinesFetcher interface
func (m *Machines) Allocate(params *maas.MachinesAllocateParams) ([]byte, error) {
	qsp := maas.ToQSP(params)
	return m.callPost("allocate", qsp)
}

// Release fulfills the  maas.MachinesFetcher interface
func (m *Machines) Release(systemIDs []string, comment string) error {
	qsp := url.Values{}
	for _, val := range systemIDs {
		qsp.Add("machines", val)
	}
	_, err := m.callPost("release", qsp)
	return err
}
