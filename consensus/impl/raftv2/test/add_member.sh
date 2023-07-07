#!/usr/bin/env bash
source test_common.sh

addnode=$1

echo "add member $addnode"

if [ "$1" = "" ] || [[ "$1" != aergo* ]] ;then
	echo "use:add_member.sh aergo4|aergo5|aergo6|aergo7 usebackup"
	exit 100
fi

usebackup=0
if [ "$2" == "usebackup" ]; then
	usebackup=1	

	echo "try to join $addnode with backup"
fi

reqport=$3
if [ "$reqport" = "" ];then
	reqport=10001
fi

leader=$(aergocli -p $reqport blockchain | jq .ConsensusInfo.Status.Leader)
leader=${leader//\"/}
if [[ $leader != aergo* ]]; then
	echo "leader not exist"
	exit 100
fi

leaderport=${ports[$leader]}
prevCnt=$(getClusterTotal $leaderport)

echo "leader=$leader, port=$leaderport, prevTotal=$prevCnt"

# By RPC
#echo "aergocli -p $leaderport cluster add --name \"$addnode\" --address \"/ip4/127.0.0.1/${httpports[$addnode]}\" --peerid \"${peerids[$addnode]}\""
#aergocli -p $leaderport cluster add --name "$addnode" --address "/ip4/127.0.0.1/${httpports[$addnode]}" --peerid "${peerids[$addnode]}"

# By Enterprise Tx
walletFile="$TEST_RAFT_INSTANCE/genesis_wallet.txt"
ADMIN=
getAdminUnlocked $leaderport $walletFile ADMIN

addJson="$(makeAddMemberJson $addnode)"

echo "aergocli -p "$leaderport" contract call --governance "$ADMIN" aergo.enterprise changeCluster "$addJson" --keystore . --password 1234"
aergocli -p $leaderport contract call --governance $ADMIN aergo.enterprise changeCluster "$addJson" --keystore . --password 1234
echo "add Done" 

mySvrport=${svrports[$addnode]}
mySvrName=${svrname[$addnode]}
myConfig="$mySvrName.toml"


for i in {1..5}; do
	echo ${svrname["aergo$i"]}
done
echo "new svr=$mySvrport $mySvrName, $myConfig"

reqCnt=$((prevCnt+1))
echo "waitClusterTotal=$reqCnt"
waitClusterTotal $reqCnt 600 $reqport
if [ $? -ne 1 ]; then
	echo "add failed"
	exit 100
fi

echo "Get config of added member"
echo "cp $TEST_RAFT_INSTANCE_CONF/$myConfig $TEST_RAFT_INSTANCE"
cp $TEST_RAFT_INSTANCE_CONF/$myConfig $TEST_RAFT_INSTANCE

pushd $TEST_RAFT_INSTANCE
if [ "$usebackup" == "0" ]; then
	echo "init genesis for $mySvrName"
	init_genesis.sh $mySvrName > /dev/null 2>&1
else 
	echo "join using backup: $mySvrName"
	do_sed.sh $myConfig "usebackup=false" "usebackup=true" "/"
	run_svr.sh $mySvrName
fi
popd
