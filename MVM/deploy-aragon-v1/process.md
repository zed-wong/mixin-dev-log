# 1. Add `mvm` in network-config.js

```
const MVM_INFORMATION = {
  id: 73927,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  type: 'mvm',
  fullName: 'Mixin Virtual Machine',
  shortName: 'MVM',
  explorerUrl: `https://scan.mvm.dev`,
  testnet: false,
}

['mvm']: {
    isActive: true,
    addresses: {
      ensRegistry:
        localEnsRegistryAddress || '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
      governExecutorProxy: null,
    },
    nodes: {
      defaultEth: 'wss://mainnet.eth.aragon.network/ws',
    },
    connectGraphEndpoint: null,
    settings: {
      chainId: 73927,
      testnet: false,
      ...MVM_INFORMATION,
    },
  },
```

2. 

Edit:

- /client/src/templates/dandelion/config/helpers/tokens.js
Line 14: ETH

- /client/src/onboarding/Onboarding/Onboarding.js
Line 101: ETH
Line 120: ETH

- /client/src/aragonjs-wrapper.js
Line 296: 



x. deploy ens registry


