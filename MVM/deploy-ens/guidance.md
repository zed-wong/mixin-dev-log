Follow:https://docs.ens.domains/deploying-ens-on-a-private-chain

Basically:
1. Create a project using hardhat
2. Import ens packages
3. Deploy the registry
4. Deploy the resolver
5. Deploy a register
6. Deploy the reverse register


What I did:
1. Deployed registry and resolver with aragon script
https://github.com/aragon/osx/blob/39ffa73f91454f75bab6920c4bec2dc8a0da56eb/packages/contracts/deploy/new/10_framework/00_ens_registry.ts#L26-L28

On MVM
ENS Registry: 0x46c43421F7f25221917D3CC4A41d7E2E34D6A0d7
ENS Resolver: 0x75078E2198733ee223B9AE3338A5b8E7aF613E22
