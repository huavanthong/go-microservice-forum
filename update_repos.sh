#!/bin/bash

# Khai báo giá trị mới cho REPOSITORY và PULL_MERGE_BRANCH

# Đường dẫn đến tệp repository.conf
CONFIG_FILE="./config/repositories.conf"
echo "Check repos: $REPOSITORY"
echo "Check branch: $PULL_MERGE_BRANCH"
echo "Check version: $TEST_VERSION"
echo "Check date: $TEST_DATE"
echo "Check at env - help me: $HELP_ME"

package_context='{
    "test_version"          : "'"${TEST_VERSION}"'",
    "test_date"             : "'"${TEST_DATE}"'"
}'

test_version=$(echo "$package_context" | jq -r '.test_version')
test_date=$(echo "$package_context" | jq -r '.test_date')

echo "Test Version: $test_version"
echo "Test Date: $test_date"


if [ -n "$REPOSITORY"  ] && [ -n "$PULL_MERGE_BRANCH" ]; then 
    echo "Exist to update"
     # Thay đổi giá trị trong tệp repository.conf
    sed -i "s/$REPOSITORY=.*/$REPOSITORY=$PULL_MERGE_BRANCH" $CONFIG_FILE
    echo "Updated $REPOSITORY in $CONFIG_FILE with value $PULL_MERGE_BRANCH"
fi



# sed -i 's/microservice-golang=\(.*\)/microservice-golang=test/' repositories.conf

# # Kiểm tra xem tệp tồn tại trước khi thay đổi
# if [ -f "$CONFIG_FILE" ]; then
#     # Thay đổi giá trị trong tệp repository.conf
#     sed -i "s/$REPOSITORY=\(.*\)/${REPOSITORY}=$PULL_MERGE_BRANCH/" "$CONFIG_FILE"
#     echo "Updated $REPOSITORY in $CONFIG_FILE with value $PULL_MERGE_BRANCH"
# else
#     echo "$CONFIG_FILE does not exist."
# fi
