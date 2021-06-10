package main

import (
	"Exper"
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"io"
	"os"
	"time"

	//"syndtr/goleveldb/leveldb"
)
var(
	key []byte
	value []byte
	WriteTime float64
	temp []byte
	t3 time.Time
	t4 time.Time
	t float64
)
// 对 v   于EthData-202w，总条目64894200，键值对数目为32447100，长度为32字节的Key有2343907，确实是1/10阿
// 所有数据字节为4715961535，约4.4G左右，数据比较正确科学。
// E:/record/address.txt 中有一百万个账户，我们用六七十万即可

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func main() {
	idle0, total0 := Exper.GetCPUSample()
	db, err := leveldb.OpenFile("/media/czh/sn/DB2/Init_80G", &opt.Options{
		Compression:opt.NoCompression, // none,不启用snappy压缩机制，
		WriteBuffer:128* opt.MiB,
		//DisableBlockCache:true,
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fi, err := os.Open("/media/czh/sn/ExpData")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//defer db2.Close()Double
	defer fi.Close()
	size:=0
	size32:=0
	sizeNo32:=0
	totali:=0
	keyi:=0
	br := bufio.NewReader(fi)
	t1:=time.Now()
	MPT:=0
	for {
		if totali%2 == 0{
			a,  c := br.ReadString('\n')
			key, _ =hex.DecodeString(a)
			//size += len(key)
			if len(key) == 32{
				MPT++
			}
			keyi++
			if c == io.EOF {
				break
			}
		}else{
			a,  c := br.ReadString('\n')
			value, _ =hex.DecodeString(a)
			//size += len(value)
			if c == io.EOF {
				break
			}
			if len(key) == 32{
				t3=time.Now()
				_ = db.Put(key, value, nil)
				t4=time.Now()
				t=t4.Sub(t3).Seconds()
				//_ = db2.Put(key, value, nil)
				size32 += len(key)+len(value)
			}else {
					t3=time.Now()
					_ = db.Put(key, value, nil)
					t4=time.Now()
					t=t4.Sub(t3).Seconds()
					//_ = db2.Put_s(key, value, nil)
					sizeNo32 += len(key)+len(value)
			}
			WriteTime += t
		}
		totali++
	}
	size = size32 + sizeNo32
	t2:=time.Now()
	db.PrintTime()
	fmt.Println("发生次数:",leveldb.Count)
	fmt.Println("总时间",t2.Sub(t1).Seconds(),"Put时间",WriteTime)
	fmt.Println("总条目数:",totali,"key数目:",keyi,"MPT数目:",MPT)
	fmt.Println("总大小:",size,size32,sizeNo32,size32+sizeNo32)
	f := float64(size / 1024 / 1024)
	fmt.Println("吞吐量为:",float64(f/WriteTime),float64(keyi)/WriteTime)
	fmt.Println(db.GetmemComp())
	fmt.Println(db.Getlevel0Comp())
	fmt.Println(db.Getnonlevel0Comp())

	idle1, total1 := Exper.GetCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
}
