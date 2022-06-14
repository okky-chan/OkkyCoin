// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package okky

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OkkyTask is an auto generated low-level Go binding around an user-defined struct.
type OkkyTask struct {
	Content string
	Status  bool
}

// OkkyMetaData contains all meta data concerning the Okky contract.
var OkkyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structOkky.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structOkky.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611042806100606000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639507d39a1161005b5780639507d39a146100d8578063b0c8f9dc14610108578063f745630f14610124578063f851a440146101405761007d565b80630f560cd7146100825780634cc82215146100a0578063751ef753146100bc575b600080fd5b61008a61015e565b6040516100979190610b1f565b60405180910390f35b6100ba60048036038101906100b59190610b8b565b6102c2565b005b6100d660048036038101906100d19190610b8b565b61043d565b005b6100f260048036038101906100ed9190610b8b565b610508565b6040516100ff9190610bf5565b60405180910390f35b610122600480360381019061011d9190610d4c565b610649565b005b61013e60048036038101906101399190610d95565b610726565b005b6101486107bc565b6040516101559190610e32565b60405180910390f35b60603373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101b857600080fd5b6001805480602002602001604051908101604052809291908181526020016000905b828210156102b9578382906000526020600020906002020160405180604001604052908160008201805461020d90610e7c565b80601f016020809104026020016040519081016040528092919081815260200182805461023990610e7c565b80156102865780601f1061025b57610100808354040283529160200191610286565b820191906000526020600020905b81548152906001019060200180831161026957829003601f168201915b505050505081526020016001820160009054906101000a900460ff161515151581525050815260200190600101906101da565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461031a57600080fd5b60008190505b600180805490506103319190610edc565b8110156103ea57600180826103469190610f10565b8154811061035757610356610f66565b5b90600052602060002090600202016001828154811061037957610378610f66565b5b9060005260206000209060020201600082018160000190805461039b90610e7c565b6103a69291906107e0565b506001820160009054906101000a900460ff168160010160006101000a81548160ff02191690831515021790555090505080806103e290610f95565b915050610320565b5060018054806103fd576103fc610fdd565b5b600190038181906000526020600020906002020160008082016000610422919061086d565b6001820160006101000a81549060ff02191690555050905550565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461049557600080fd5b600181815481106104a9576104a8610f66565b5b906000526020600020906002020160010160009054906101000a900460ff1615600182815481106104dd576104dc610f66565b5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555050565b6105106108ad565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461056857600080fd5b6001828154811061057c5761057b610f66565b5b90600052602060002090600202016040518060400160405290816000820180546105a590610e7c565b80601f01602080910402602001604051908101604052809291908181526020018280546105d190610e7c565b801561061e5780601f106105f35761010080835404028352916020019161061e565b820191906000526020600020905b81548152906001019060200180831161060157829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106a157600080fd5b6001604051806040016040528083815260200160001515815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906107009291906108c9565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461077e57600080fd5b806001838154811061079357610792610f66565b5b906000526020600020906002020160000190805190602001906107b79291906108c9565b505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8280546107ec90610e7c565b90600052602060002090601f01602090048101928261080e576000855561085c565b82601f1061081f578054855561085c565b8280016001018555821561085c57600052602060002091601f016020900482015b8281111561085b578254825591600101919060010190610840565b5b509050610869919061094f565b5090565b50805461087990610e7c565b6000825580601f1061088b57506108aa565b601f0160209004906000526020600020908101906108a9919061094f565b5b50565b6040518060400160405280606081526020016000151581525090565b8280546108d590610e7c565b90600052602060002090601f0160209004810192826108f7576000855561093e565b82601f1061091057805160ff191683800117855561093e565b8280016001018555821561093e579182015b8281111561093d578251825591602001919060010190610922565b5b50905061094b919061094f565b5090565b5b80821115610968576000816000905550600101610950565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109d25780820151818401526020810190506109b7565b838111156109e1576000848401525b50505050565b6000601f19601f8301169050919050565b6000610a0382610998565b610a0d81856109a3565b9350610a1d8185602086016109b4565b610a26816109e7565b840191505092915050565b60008115159050919050565b610a4681610a31565b82525050565b60006040830160008301518482036000860152610a6982826109f8565b9150506020830151610a7e6020860182610a3d565b508091505092915050565b6000610a958383610a4c565b905092915050565b6000602082019050919050565b6000610ab58261096c565b610abf8185610977565b935083602082028501610ad185610988565b8060005b85811015610b0d5784840389528151610aee8582610a89565b9450610af983610a9d565b925060208a01995050600181019050610ad5565b50829750879550505050505092915050565b60006020820190508181036000830152610b398184610aaa565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610b6881610b55565b8114610b7357600080fd5b50565b600081359050610b8581610b5f565b92915050565b600060208284031215610ba157610ba0610b4b565b5b6000610baf84828501610b76565b91505092915050565b60006040830160008301518482036000860152610bd582826109f8565b9150506020830151610bea6020860182610a3d565b508091505092915050565b60006020820190508181036000830152610c0f8184610bb8565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c59826109e7565b810181811067ffffffffffffffff82111715610c7857610c77610c21565b5b80604052505050565b6000610c8b610b41565b9050610c978282610c50565b919050565b600067ffffffffffffffff821115610cb757610cb6610c21565b5b610cc0826109e7565b9050602081019050919050565b82818337600083830152505050565b6000610cef610cea84610c9c565b610c81565b905082815260208101848484011115610d0b57610d0a610c1c565b5b610d16848285610ccd565b509392505050565b600082601f830112610d3357610d32610c17565b5b8135610d43848260208601610cdc565b91505092915050565b600060208284031215610d6257610d61610b4b565b5b600082013567ffffffffffffffff811115610d8057610d7f610b50565b5b610d8c84828501610d1e565b91505092915050565b60008060408385031215610dac57610dab610b4b565b5b6000610dba85828601610b76565b925050602083013567ffffffffffffffff811115610ddb57610dda610b50565b5b610de785828601610d1e565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e1c82610df1565b9050919050565b610e2c81610e11565b82525050565b6000602082019050610e476000830184610e23565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e9457607f821691505b602082108103610ea757610ea6610e4d565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ee782610b55565b9150610ef283610b55565b925082821015610f0557610f04610ead565b5b828203905092915050565b6000610f1b82610b55565b9150610f2683610b55565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f5b57610f5a610ead565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610fa082610b55565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fd257610fd1610ead565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212201573eb533fd766af106e3ba65848c91e6ef4ad214b51019676021ad222b14de864736f6c634300080e0033",
}

// OkkyABI is the input ABI used to generate the binding from.
// Deprecated: Use OkkyMetaData.ABI instead.
var OkkyABI = OkkyMetaData.ABI

// OkkyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OkkyMetaData.Bin instead.
var OkkyBin = OkkyMetaData.Bin

// DeployOkky deploys a new Ethereum contract, binding an instance of Okky to it.
func DeployOkky(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Okky, error) {
	parsed, err := OkkyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OkkyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Okky{OkkyCaller: OkkyCaller{contract: contract}, OkkyTransactor: OkkyTransactor{contract: contract}, OkkyFilterer: OkkyFilterer{contract: contract}}, nil
}

// Okky is an auto generated Go binding around an Ethereum contract.
type Okky struct {
	OkkyCaller     // Read-only binding to the contract
	OkkyTransactor // Write-only binding to the contract
	OkkyFilterer   // Log filterer for contract events
}

// OkkyCaller is an auto generated read-only Go binding around an Ethereum contract.
type OkkyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkkyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OkkyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkkyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OkkyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkkySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OkkySession struct {
	Contract     *Okky             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OkkyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OkkyCallerSession struct {
	Contract *OkkyCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OkkyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OkkyTransactorSession struct {
	Contract     *OkkyTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OkkyRaw is an auto generated low-level Go binding around an Ethereum contract.
type OkkyRaw struct {
	Contract *Okky // Generic contract binding to access the raw methods on
}

// OkkyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OkkyCallerRaw struct {
	Contract *OkkyCaller // Generic read-only contract binding to access the raw methods on
}

// OkkyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OkkyTransactorRaw struct {
	Contract *OkkyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOkky creates a new instance of Okky, bound to a specific deployed contract.
func NewOkky(address common.Address, backend bind.ContractBackend) (*Okky, error) {
	contract, err := bindOkky(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Okky{OkkyCaller: OkkyCaller{contract: contract}, OkkyTransactor: OkkyTransactor{contract: contract}, OkkyFilterer: OkkyFilterer{contract: contract}}, nil
}

// NewOkkyCaller creates a new read-only instance of Okky, bound to a specific deployed contract.
func NewOkkyCaller(address common.Address, caller bind.ContractCaller) (*OkkyCaller, error) {
	contract, err := bindOkky(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OkkyCaller{contract: contract}, nil
}

// NewOkkyTransactor creates a new write-only instance of Okky, bound to a specific deployed contract.
func NewOkkyTransactor(address common.Address, transactor bind.ContractTransactor) (*OkkyTransactor, error) {
	contract, err := bindOkky(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OkkyTransactor{contract: contract}, nil
}

// NewOkkyFilterer creates a new log filterer instance of Okky, bound to a specific deployed contract.
func NewOkkyFilterer(address common.Address, filterer bind.ContractFilterer) (*OkkyFilterer, error) {
	contract, err := bindOkky(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OkkyFilterer{contract: contract}, nil
}

// bindOkky binds a generic wrapper to an already deployed contract.
func bindOkky(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OkkyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Okky *OkkyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Okky.Contract.OkkyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Okky *OkkyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Okky.Contract.OkkyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Okky *OkkyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Okky.Contract.OkkyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Okky *OkkyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Okky.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Okky *OkkyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Okky.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Okky *OkkyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Okky.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Okky *OkkyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Okky.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Okky *OkkySession) Admin() (common.Address, error) {
	return _Okky.Contract.Admin(&_Okky.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Okky *OkkyCallerSession) Admin() (common.Address, error) {
	return _Okky.Contract.Admin(&_Okky.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Okky *OkkyCaller) Get(opts *bind.CallOpts, _id *big.Int) (OkkyTask, error) {
	var out []interface{}
	err := _Okky.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(OkkyTask), err
	}

	out0 := *abi.ConvertType(out[0], new(OkkyTask)).(*OkkyTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Okky *OkkySession) Get(_id *big.Int) (OkkyTask, error) {
	return _Okky.Contract.Get(&_Okky.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Okky *OkkyCallerSession) Get(_id *big.Int) (OkkyTask, error) {
	return _Okky.Contract.Get(&_Okky.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Okky *OkkyCaller) List(opts *bind.CallOpts) ([]OkkyTask, error) {
	var out []interface{}
	err := _Okky.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]OkkyTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]OkkyTask)).(*[]OkkyTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Okky *OkkySession) List() ([]OkkyTask, error) {
	return _Okky.Contract.List(&_Okky.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Okky *OkkyCallerSession) List() ([]OkkyTask, error) {
	return _Okky.Contract.List(&_Okky.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Okky *OkkyTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Okky.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Okky *OkkySession) Add(_content string) (*types.Transaction, error) {
	return _Okky.Contract.Add(&_Okky.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Okky *OkkyTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Okky.Contract.Add(&_Okky.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Okky *OkkyTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Okky.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Okky *OkkySession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Okky.Contract.Remove(&_Okky.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Okky *OkkyTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Okky.Contract.Remove(&_Okky.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Okky *OkkyTransactor) Toggle(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Okky.contract.Transact(opts, "toggle", _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Okky *OkkySession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Okky.Contract.Toggle(&_Okky.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Okky *OkkyTransactorSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Okky.Contract.Toggle(&_Okky.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Okky *OkkyTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Okky.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Okky *OkkySession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Okky.Contract.Update(&_Okky.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Okky *OkkyTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Okky.Contract.Update(&_Okky.TransactOpts, _id, _content)
}
