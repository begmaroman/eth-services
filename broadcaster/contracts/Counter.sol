pragma solidity ^0.8.0;

contract Counter {
   event Performed(
      address indexed from,
      uint256 initialBlock,
      uint256 lastBlock,
      uint256 previousBlock,
      uint256 counter
   );

   event TriggerPerformance(
      bytes data,
      bool perform
   );

   uint256 public testRange;
   uint256 public interval;
   uint256 public lastBlock;
   uint256 public previousPerformBlock;
   uint256 public initialBlock;
   uint256 public counter;

   constructor(uint256 _testRange, uint256 _interval) {
      testRange = _testRange;
      interval = _interval;
      previousPerformBlock = 0;
      lastBlock = block.number;
      initialBlock = 0;
      counter = 0;
   }

   function trigger(bytes calldata data, bool perform) external {
      emit TriggerPerformance(data, perform);
   }

   function check(bytes calldata data) external view returns (bool, bytes memory) {
      return (eligible(), data);
   }

   function perform(bytes calldata performData) external {
      if (initialBlock == 0) {
         initialBlock = block.number;
      }
      lastBlock = block.number;
      counter = counter + 1;
      performData;
      emit Performed(tx.origin, initialBlock, lastBlock, previousPerformBlock, counter);
      previousPerformBlock = lastBlock;
   }

   function eligible() public view returns (bool) {
      if (initialBlock == 0) {
         return true;
      }

      return (block.number - initialBlock) < testRange && (block.number - lastBlock) >= interval;
   }

   function setSpread(uint256 _testRange, uint256 _interval) external {
      testRange = _testRange;
      interval = _interval;
      initialBlock = 0;
      counter = 0;
   }
}
