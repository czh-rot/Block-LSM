package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main()  {
	db, err := leveldb.OpenFile("/media/czh/sn/DB2/Test_1", &opt.Options{
		Compression:opt.NoCompression, // none,不启用snappy压缩机制，
		WriteBuffer:1* opt.MiB,
		//DisableBlockCache:true,
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Put([]byte{10,1,1,1,1},[]byte{2},nil)
	db.Put([]byte{11,1,1,1,2},[]byte{2},nil)
	//db.Put([]byte{12,1,1,1,3},[]byte{2},nil)
	//db.Put([]byte{13,1,1,1,4},[]byte{2},nil)
	//db.Put([]byte{14,1,1,1,5},[]byte{2},nil)
	//db.Put([]byte{15,1,1,1,6},[]byte{2},nil)
	//db.Put([]byte{15,1,1,1,7},[]byte{2},nil)
	//db.Put([]byte{17,1,1,1,8},[]byte{2},nil)
	//db.Put([]byte{18,1,1,1,9},[]byte{2},nil)
	//db.Put([]byte{19,1,1,1,10},[]byte{2},nil)
	//db.Put([]byte{20,1,1,1,11},[]byte{2},nil)
}
