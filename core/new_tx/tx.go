package tx

import (
	"github.com/iost-official/Go-IOS-Protocol/common"
	"time"
	"github.com/iost-official/Go-IOS-Protocol/account"
	"github.com/gogo/protobuf/proto"
)

//go:generate gencode go -schema=structs.schema -package=tx

// Tx Transaction 的实现
type Tx struct {
	// TODO calculate id
	Id			string					// encode tx hash
	Time      	int64
	Actions		[]Action
	Signers		[][]byte
	Signs     	[]common.Signature
	Publisher 	common.Signature
}

// 新建一个Tx，需要通过编译器得到一个contract
func NewTx(nonce int64, actions []Action, signers [][]byte) Tx {
	return Tx{
		Time:     	time.Now().UnixNano(),
		Actions:	actions,
		Signers:	signers,
	}
}

// 合约的参与者进行签名
func SignTxContent(tx Tx, account account.Account) (common.Signature, error) {
	sign, err := common.Sign(common.Secp256k1, tx.baseHash(), account.Seckey)
	if err != nil {
		return sign, err
	}
	return sign, nil
}
// Time,Noce,Contract形成的基本哈希值
func (t *Tx) baseHash() []byte {
	tr := &TxRaw{
		Id:t.Id,
		Time:t.Time,
	}
	for _, a := range t.Actions {
		tr.Actions = append(tr.Actions, &ActionRaw{
			Contract:a.Contract,
			ActionName:a.ActionName,
			Data:a.Data,
		})
	}
	tr.Signers = t.Signers

	b, err := proto.Marshal(tr)
	if err != nil {
		panic(err)
	}
	return common.Sha256(b)
}

// 合约的发布者进行签名，此签名的用户用于支付gas
func SignTx(tx Tx, account account.Account, signs ...common.Signature) (Tx, error) {
	tx.Signs = append(tx.Signs, signs...)
	sign, err := common.Sign(common.Secp256k1, tx.publishHash(), account.Seckey)
	if err != nil {
		return tx, err
	}
	tx.Publisher = sign
	return tx, nil
}


// publishHash 发布者使用的hash值，包含参与者的签名
func (t *Tx) publishHash() []byte {
	tr := &TxRaw{
		Id:t.Id,
		Time:t.Time,
	}
	for _, a := range t.Actions {
		tr.Actions = append(tr.Actions, &ActionRaw{
			Contract:a.Contract,
			ActionName:a.ActionName,
			Data:a.Data,
		})
	}
	tr.Signers = t.Signers
	for _, s := range t.Signs {
		tr.Signs = append(tr.Signs, &common.SignatureRaw{
			Algorithm:int32(s.Algorithm),
			Sig:s.Sig,
			PubKey:s.Pubkey,
		})
	}

	b, err := proto.Marshal(tr)
	if err != nil {
		panic(err)
	}
	return common.Sha256(b)
}

// 对Tx进行编码
func (t *Tx) Encode() []byte {
	tr := &TxRaw{
		Id:t.Id,
		Time:t.Time,
	}
	for _, a := range t.Actions {
		tr.Actions = append(tr.Actions, &ActionRaw{
			Contract:a.Contract,
			ActionName:a.ActionName,
			Data:a.Data,
		})
	}
	tr.Signers = t.Signers
	for _, s := range t.Signs {
		tr.Signs = append(tr.Signs, &common.SignatureRaw{
			Algorithm:int32(s.Algorithm),
			Sig:s.Sig,
			PubKey:s.Pubkey,
		})
	}
	tr.Publisher = &common.SignatureRaw{
		Algorithm:int32(t.Publisher.Algorithm),
		Sig:t.Publisher.Sig,
		PubKey:t.Publisher.Pubkey,
	}

	b, err := proto.Marshal(tr)
	if err != nil {
		panic(err)
	}
	return b
}

// 对Tx进行解码
func (t *Tx) Decode(b []byte) error {
	var tr *TxRaw
	err := proto.Unmarshal(b, tr)
	if err != nil {
		return err
	}
	t.Id = tr.Id
	t.Time = tr.Time
	t.Actions = []Action{}
	for _, a := range tr.Actions {
		t.Actions = append(t.Actions, Action{
			Contract:a.Contract,
			ActionName:a.ActionName,
			Data:a.Data,
		})
	}
	t.Signers = tr.Signers
	t.Signs = []common.Signature{}
	for _, sr := range tr.Signs {
		t.Signs = append(t.Signs, common.Signature{
			Algorithm:common.SignAlgorithm(sr.Algorithm),
			Sig:sr.Sig,
			Pubkey:sr.PubKey,
		})
	}
	t.Publisher = common.Signature{
		Algorithm: common.SignAlgorithm(tr.Publisher.Algorithm),
		Sig:tr.Publisher.Sig,
		Pubkey:tr.Publisher.PubKey,
	}

	return nil
}

// 计算Tx的哈希值
func (t *Tx) Hash() []byte {
	return common.Sha256(t.Encode())
}

/*
// 验证签名的函数
func (t *Tx) VerifySelf() error {
	baseHash := t.baseHash() // todo 在basehash内缓存，不需要在应用进行缓存
	for _, sign := range t.Signs {
		ok := common.VerifySignature(baseHash, sign)
		if !ok {
			return fmt.Errorf("signer error")
		}
	}

	ok := common.VerifySignature(t.publishHash(), t.Publisher)
	if !ok {
		return fmt.Errorf("publisher error")
	}
	return nil
}

func (t *Tx) VerifySigner(sig common.Signature) bool {
	return common.VerifySignature(t.baseHash(), sig)
}
*/