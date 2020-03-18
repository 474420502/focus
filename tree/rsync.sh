
rsync() {
    FILE_NAME=$3  #iterator_test.go
    SRC_DIR_NAME=$1 #vbtkey dir
    DEST_DIR_NAME=$2 # vbtdupkey dir

    cp ./$SRC_DIR_NAME/$FILE_NAME ./$DEST_DIR_NAME/$FILE_NAME
    sed -i "s/package $SRC_DIR_NAME/package ${DEST_DIR_NAME}/" ./$DEST_DIR_NAME/$FILE_NAME
}

# 同步iterator_test
for dst in "vbtkeydup" "avlkey" "avlkeydup" 
do
    rsync "vbtkey" $dst "iterator_test.go"
done

# 同步iterator_test
for dst in "vbtdup" "avl" "avldup" 
do
    rsync "vbt" $dst "iterator_test.go"
done
