//SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;
pragma abicoder v2;

contract Okky {

    address public admin;
    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        admin = msg.sender;
    }

    modifier isAdmin() {
        require(admin == msg.sender);
        _;
    }

    function add(string memory _content) public isAdmin {
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public isAdmin view returns(Task memory) {
        return tasks[_id];
    }

    function list() public isAdmin view returns(Task[] memory) {
        return tasks;
    }

    function update(uint _id, string memory _content) public isAdmin {
        tasks[_id].content = _content;
    }

    function toggle(uint _id) public isAdmin {
        tasks[_id].status = !tasks[_id].status;
    }

    function remove(uint _id) public isAdmin {
        for(uint i = _id; i < tasks.length - 1; i++) {
            tasks[i] = tasks[i + 1];
        }
        tasks.pop();
    }
}