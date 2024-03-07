// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 < 0.9.0;

contract Todo {
 
 Task[] tasks ;
 address public owner;

     struct Task{

        string content;
        bool status;
            
        }

    constructor (){
           owner = msg.sender;
       
    }
    modifier isOwner(){
            require(msg.sender == owner);
            _;
    }


    function add(string memory _content) public isOwner {
        

        tasks.push(Task(_content, false));
    }
    function get(uint id) public isOwner view  returns (Task memory) {

        return  tasks[id];
    }
    function list() public  isOwner view  returns (Task[] memory)  {
        return tasks;
    }
    function update(uint _id , string memory _content) public isOwner {

        tasks[_id].content = _content;
            
    }
    function remove(uint _id) public  isOwner {

        for ( uint i= _id; i < tasks.length - 1; i++) 
        {
            tasks[i] = tasks[i+1];
        }
        tasks.pop();
    }

    
   
}