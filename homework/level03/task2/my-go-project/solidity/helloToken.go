// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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
	_ = abi.ConvertType
)

// HelloTokenMetaData contains all meta data concerning the HelloToken contract.
var HelloTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611881380380611881833981810160405281019061003291906103eb565b6040518060400160405280600a81526020017f48656c6c6f546f6b656e000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f48544e000000000000000000000000000000000000000000000000000000000081525081600390816100ad919061065e565b5080600490816100bd919061065e565b5050506100f4336100d26100fa60201b60201c565b600a6100de919061089f565b836100e991906108ea565b61010360201b60201c565b50610a1d565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101755760006040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161016c919061096d565b60405180910390fd5b6101876000838361018b60201b60201c565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036101dd5780600260008282546101d19190610988565b925050819055506102b0565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610269578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610260939291906109cb565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102f95780600260008282540392505081905550610346565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103a39190610a02565b60405180910390a3505050565b600080fd5b6000819050919050565b6103c8816103b5565b81146103d357600080fd5b50565b6000815190506103e5816103bf565b92915050565b600060208284031215610401576104006103b0565b5b600061040f848285016103d6565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061049957607f821691505b6020821081036104ac576104ab610452565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026105147fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826104d7565b61051e86836104d7565b95508019841693508086168417925050509392505050565b6000819050919050565b600061055b610556610551846103b5565b610536565b6103b5565b9050919050565b6000819050919050565b61057583610540565b61058961058182610562565b8484546104e4565b825550505050565b600090565b61059e610591565b6105a981848461056c565b505050565b5b818110156105cd576105c2600082610596565b6001810190506105af565b5050565b601f821115610612576105e3816104b2565b6105ec846104c7565b810160208510156105fb578190505b61060f610607856104c7565b8301826105ae565b50505b505050565b600082821c905092915050565b600061063560001984600802610617565b1980831691505092915050565b600061064e8383610624565b9150826002028217905092915050565b61066782610418565b67ffffffffffffffff8111156106805761067f610423565b5b61068a8254610481565b6106958282856105d1565b600060209050601f8311600181146106c857600084156106b6578287015190505b6106c08582610642565b865550610728565b601f1984166106d6866104b2565b60005b828110156106fe578489015182556001820191506020850194506020810190506106d9565b8683101561071b5784890151610717601f891682610624565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156107b65780860481111561079257610791610730565b5b60018516156107a15780820291505b80810290506107af8561075f565b9450610776565b94509492505050565b6000826107cf576001905061088b565b816107dd576000905061088b565b81600181146107f357600281146107fd5761082c565b600191505061088b565b60ff84111561080f5761080e610730565b5b8360020a91508482111561082657610825610730565b5b5061088b565b5060208310610133831016604e8410600b84101617156108615782820a90508381111561085c5761085b610730565b5b61088b565b61086e848484600161076c565b9250905081840481111561088557610884610730565b5b81810290505b9392505050565b600060ff82169050919050565b60006108aa826103b5565b91506108b583610892565b92506108e27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846107bf565b905092915050565b60006108f5826103b5565b9150610900836103b5565b925082820261090e816103b5565b9150828204841483151761092557610924610730565b5b5092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109578261092c565b9050919050565b6109678161094c565b82525050565b6000602082019050610982600083018461095e565b92915050565b6000610993826103b5565b915061099e836103b5565b92508282019050808211156109b6576109b5610730565b5b92915050565b6109c5816103b5565b82525050565b60006060820190506109e0600083018661095e565b6109ed60208301856109bc565b6109fa60408301846109bc565b949350505050565b6000602082019050610a1760008301846109bc565b92915050565b610e5580610a2c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063313ce56711610066578063313ce5671461013457806370a082311461015257806395d89b4114610182578063a9059cbb146101a0578063dd62ed3e146101d057610093565b806306fdde0314610098578063095ea7b3146100b657806318160ddd146100e657806323b872dd14610104575b600080fd5b6100a0610200565b6040516100ad9190610aa9565b60405180910390f35b6100d060048036038101906100cb9190610b64565b610292565b6040516100dd9190610bbf565b60405180910390f35b6100ee6102b5565b6040516100fb9190610be9565b60405180910390f35b61011e60048036038101906101199190610c04565b6102bf565b60405161012b9190610bbf565b60405180910390f35b61013c6102ee565b6040516101499190610c73565b60405180910390f35b61016c60048036038101906101679190610c8e565b6102f7565b6040516101799190610be9565b60405180910390f35b61018a61033f565b6040516101979190610aa9565b60405180910390f35b6101ba60048036038101906101b59190610b64565b6103d1565b6040516101c79190610bbf565b60405180910390f35b6101ea60048036038101906101e59190610cbb565b6103f4565b6040516101f79190610be9565b60405180910390f35b60606003805461020f90610d2a565b80601f016020809104026020016040519081016040528092919081815260200182805461023b90610d2a565b80156102885780601f1061025d57610100808354040283529160200191610288565b820191906000526020600020905b81548152906001019060200180831161026b57829003601f168201915b5050505050905090565b60008061029d61047b565b90506102aa818585610483565b600191505092915050565b6000600254905090565b6000806102ca61047b565b90506102d7858285610495565b6102e2858585610529565b60019150509392505050565b60006012905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606004805461034e90610d2a565b80601f016020809104026020016040519081016040528092919081815260200182805461037a90610d2a565b80156103c75780601f1061039c576101008083540402835291602001916103c7565b820191906000526020600020905b8154815290600101906020018083116103aa57829003601f168201915b5050505050905090565b6000806103dc61047b565b90506103e9818585610529565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b610490838383600161061d565b505050565b60006104a184846103f4565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146105235781811015610513578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161050a93929190610d6a565b60405180910390fd5b6105228484848403600061061d565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361059b5760006040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016105929190610da1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361060d5760006040517fec442f050000000000000000000000000000000000000000000000000000000081526004016106049190610da1565b60405180910390fd5b6106188383836107f4565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160361068f5760006040517fe602df050000000000000000000000000000000000000000000000000000000081526004016106869190610da1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107015760006040517f94280d620000000000000000000000000000000000000000000000000000000081526004016106f89190610da1565b60405180910390fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080156107ee578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107e59190610be9565b60405180910390a35b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361084657806002600082825461083a9190610deb565b92505081905550610919565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156108d2578381836040517fe450d38c0000000000000000000000000000000000000000000000000000000081526004016108c993929190610d6a565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361096257806002600082825403925050819055506109af565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610a0c9190610be9565b60405180910390a3505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a53578082015181840152602081019050610a38565b60008484015250505050565b6000601f19601f8301169050919050565b6000610a7b82610a19565b610a858185610a24565b9350610a95818560208601610a35565b610a9e81610a5f565b840191505092915050565b60006020820190508181036000830152610ac38184610a70565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610afb82610ad0565b9050919050565b610b0b81610af0565b8114610b1657600080fd5b50565b600081359050610b2881610b02565b92915050565b6000819050919050565b610b4181610b2e565b8114610b4c57600080fd5b50565b600081359050610b5e81610b38565b92915050565b60008060408385031215610b7b57610b7a610acb565b5b6000610b8985828601610b19565b9250506020610b9a85828601610b4f565b9150509250929050565b60008115159050919050565b610bb981610ba4565b82525050565b6000602082019050610bd46000830184610bb0565b92915050565b610be381610b2e565b82525050565b6000602082019050610bfe6000830184610bda565b92915050565b600080600060608486031215610c1d57610c1c610acb565b5b6000610c2b86828701610b19565b9350506020610c3c86828701610b19565b9250506040610c4d86828701610b4f565b9150509250925092565b600060ff82169050919050565b610c6d81610c57565b82525050565b6000602082019050610c886000830184610c64565b92915050565b600060208284031215610ca457610ca3610acb565b5b6000610cb284828501610b19565b91505092915050565b60008060408385031215610cd257610cd1610acb565b5b6000610ce085828601610b19565b9250506020610cf185828601610b19565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610d4257607f821691505b602082108103610d5557610d54610cfb565b5b50919050565b610d6481610af0565b82525050565b6000606082019050610d7f6000830186610d5b565b610d8c6020830185610bda565b610d996040830184610bda565b949350505050565b6000602082019050610db66000830184610d5b565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610df682610b2e565b9150610e0183610b2e565b9250828201905080821115610e1957610e18610dbc565b5b9291505056fea26469706673582212206d5390a4a42b270900cc5694b7a0b2682a9376c30e513b9a941af52f0487aac364736f6c634300081c0033",
}

// HelloTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloTokenMetaData.ABI instead.
var HelloTokenABI = HelloTokenMetaData.ABI

// HelloTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloTokenMetaData.Bin instead.
var HelloTokenBin = HelloTokenMetaData.Bin

// DeployHelloToken deploys a new Ethereum contract, binding an instance of HelloToken to it.
func DeployHelloToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *HelloToken, error) {
	parsed, err := HelloTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloTokenBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HelloToken{HelloTokenCaller: HelloTokenCaller{contract: contract}, HelloTokenTransactor: HelloTokenTransactor{contract: contract}, HelloTokenFilterer: HelloTokenFilterer{contract: contract}}, nil
}

// HelloToken is an auto generated Go binding around an Ethereum contract.
type HelloToken struct {
	HelloTokenCaller     // Read-only binding to the contract
	HelloTokenTransactor // Write-only binding to the contract
	HelloTokenFilterer   // Log filterer for contract events
}

// HelloTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloTokenSession struct {
	Contract     *HelloToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloTokenCallerSession struct {
	Contract *HelloTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloTokenTransactorSession struct {
	Contract     *HelloTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloTokenRaw struct {
	Contract *HelloToken // Generic contract binding to access the raw methods on
}

// HelloTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloTokenCallerRaw struct {
	Contract *HelloTokenCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloTokenTransactorRaw struct {
	Contract *HelloTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloToken creates a new instance of HelloToken, bound to a specific deployed contract.
func NewHelloToken(address common.Address, backend bind.ContractBackend) (*HelloToken, error) {
	contract, err := bindHelloToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloToken{HelloTokenCaller: HelloTokenCaller{contract: contract}, HelloTokenTransactor: HelloTokenTransactor{contract: contract}, HelloTokenFilterer: HelloTokenFilterer{contract: contract}}, nil
}

// NewHelloTokenCaller creates a new read-only instance of HelloToken, bound to a specific deployed contract.
func NewHelloTokenCaller(address common.Address, caller bind.ContractCaller) (*HelloTokenCaller, error) {
	contract, err := bindHelloToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTokenCaller{contract: contract}, nil
}

// NewHelloTokenTransactor creates a new write-only instance of HelloToken, bound to a specific deployed contract.
func NewHelloTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTokenTransactor, error) {
	contract, err := bindHelloToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTokenTransactor{contract: contract}, nil
}

// NewHelloTokenFilterer creates a new log filterer instance of HelloToken, bound to a specific deployed contract.
func NewHelloTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloTokenFilterer, error) {
	contract, err := bindHelloToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloTokenFilterer{contract: contract}, nil
}

// bindHelloToken binds a generic wrapper to an already deployed contract.
func bindHelloToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloToken *HelloTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloToken.Contract.HelloTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloToken *HelloTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloToken.Contract.HelloTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloToken *HelloTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloToken.Contract.HelloTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloToken *HelloTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloToken *HelloTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloToken *HelloTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HelloToken *HelloTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HelloToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HelloToken *HelloTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HelloToken.Contract.Allowance(&_HelloToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HelloToken *HelloTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HelloToken.Contract.Allowance(&_HelloToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HelloToken *HelloTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HelloToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HelloToken *HelloTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HelloToken.Contract.BalanceOf(&_HelloToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HelloToken *HelloTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HelloToken.Contract.BalanceOf(&_HelloToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HelloToken *HelloTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HelloToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HelloToken *HelloTokenSession) Decimals() (uint8, error) {
	return _HelloToken.Contract.Decimals(&_HelloToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HelloToken *HelloTokenCallerSession) Decimals() (uint8, error) {
	return _HelloToken.Contract.Decimals(&_HelloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HelloToken *HelloTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HelloToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HelloToken *HelloTokenSession) Name() (string, error) {
	return _HelloToken.Contract.Name(&_HelloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HelloToken *HelloTokenCallerSession) Name() (string, error) {
	return _HelloToken.Contract.Name(&_HelloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HelloToken *HelloTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HelloToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HelloToken *HelloTokenSession) Symbol() (string, error) {
	return _HelloToken.Contract.Symbol(&_HelloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HelloToken *HelloTokenCallerSession) Symbol() (string, error) {
	return _HelloToken.Contract.Symbol(&_HelloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HelloToken *HelloTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HelloToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HelloToken *HelloTokenSession) TotalSupply() (*big.Int, error) {
	return _HelloToken.Contract.TotalSupply(&_HelloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HelloToken *HelloTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _HelloToken.Contract.TotalSupply(&_HelloToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_HelloToken *HelloTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_HelloToken *HelloTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.Contract.Approve(&_HelloToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_HelloToken *HelloTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.Contract.Approve(&_HelloToken.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_HelloToken *HelloTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_HelloToken *HelloTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.Contract.Transfer(&_HelloToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_HelloToken *HelloTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.Contract.Transfer(&_HelloToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_HelloToken *HelloTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_HelloToken *HelloTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.Contract.TransferFrom(&_HelloToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_HelloToken *HelloTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HelloToken.Contract.TransferFrom(&_HelloToken.TransactOpts, from, to, value)
}

// HelloTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HelloToken contract.
type HelloTokenApprovalIterator struct {
	Event *HelloTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HelloTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HelloTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HelloTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloTokenApproval represents a Approval event raised by the HelloToken contract.
type HelloTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HelloToken *HelloTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HelloTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HelloToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HelloTokenApprovalIterator{contract: _HelloToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HelloToken *HelloTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HelloTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HelloToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloTokenApproval)
				if err := _HelloToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HelloToken *HelloTokenFilterer) ParseApproval(log types.Log) (*HelloTokenApproval, error) {
	event := new(HelloTokenApproval)
	if err := _HelloToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HelloToken contract.
type HelloTokenTransferIterator struct {
	Event *HelloTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HelloTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HelloTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HelloTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloTokenTransfer represents a Transfer event raised by the HelloToken contract.
type HelloTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HelloToken *HelloTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HelloTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HelloToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HelloTokenTransferIterator{contract: _HelloToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HelloToken *HelloTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HelloTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HelloToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloTokenTransfer)
				if err := _HelloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HelloToken *HelloTokenFilterer) ParseTransfer(log types.Log) (*HelloTokenTransfer, error) {
	event := new(HelloTokenTransfer)
	if err := _HelloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
