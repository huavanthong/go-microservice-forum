#!/bin/bash

# Common function
function errormsg(){
   echo -e "${COL_RED}>>>> ERROR: $1!${COL_RESET}"
   echo
   exit 1
}

function goodmsg(){
   echo -e "${COL_GREEN}>>>> $1.${COL_RESET}"
   echo
}

function greenmsg(){
   echo -e "${COL_GREEN}> $1.${COL_RESET}"   
}

# Khai báo giá trị mới cho REPOSITORY và PULL_REQUEST_BRANCH

# Đường dẫn đến tệp repository.conf
config_file="./config/repositories.conf"

echo "Check repos: $REPOSITORY"
echo "Check branch: $PULL_REQUEST_BRANCH"
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


if [ -n "$REPOSITORY"  ] && [ -n "$PULL_REQUEST_BRANCH" ]; then 
	echo -e "${COL_GREEN}#          Trigger AIO from repository: $REPOSITORY; branch: $PULL_REQUEST_BRANCH    #${COL_RESET}"

	grep -q $REPOSITORY $config_file
	if [ $? -ne 0 ]; then
    	errormsg "Repository received from request not found in '$config_file'"
	else
		sed -i "s,$REPOSITORY=.*,$REPOSITORY=$PULL_REQUEST_BRANCH," $config_file
		echo -e "${COL_GREEN}#          Updated $REPOSITORY in $config_file with value $PULL_REQUEST_BRANCH"
	fi
fi



# sed -i 's/microservice-golang=\(.*\)/microservice-golang=test/' repositories.conf

# # Kiểm tra xem tệp tồn tại trước khi thay đổi
# if [ -f "$config_file" ]; then
#     # Thay đổi giá trị trong tệp repository.conf
#     sed -i "s/$REPOSITORY=\(.*\)/${REPOSITORY}=$PULL_REQUEST_BRANCH/" "$config_file"
#     echo "Updated $REPOSITORY in $config_file with value $PULL_REQUEST_BRANCH"
# else
#     echo "$config_file does not exist."
# fi
