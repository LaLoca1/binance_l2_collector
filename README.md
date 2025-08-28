# 📊 Binance L2 Order Book Collector

A lightweight Go service that connects to Binance's WebSocket API and streams **Level 2 (L2) depth data** (order book updates) into **Redis** for fast, structured access.  

The order book is a real-time list of buy (bids) and sell (asks) orders for a trading pair like BTC/USDT.
Level 2 (L2) data provides detailed info about multiple price levels, not just the best bid/ask.

Traders, market makers, and algorithms use L2 data to detect supply/demand imbalances and liquidity patterns.

---

## Financial Concepts

### 1. **Market Data Basics**

| Term | What it means |
| --- | --- |
| **Level 1 data** | Best bid/ask prices and last trade. |
| **Level 2 data** | Full order book: multiple price levels and quantities for bids & asks. |
| **Order book** | A real-time list of buy/sell orders sorted by price. |
| **Bid** | Price someone is willing to buy at. |
| **Ask** | Price someone is willing to sell at. |
| **Depth** | How many levels of bids/asks are available. |

You are collecting **Level 2 (L2)** data — which gives you **full market depth**, not just the top-of-book.

When collecting the L2 data, a **stream of JSON messages** appears like this:

{
  "e": "depthUpdate",    // event type
  "E": 123456789,        // event time
  "s": "BTCUSDT",        // symbol
  "U": 157,              // first update ID
  "u": 160,              // final update ID
  "b": [                 // bid updates
    ["29000.00", "1.2"], // price, quantity
    ...
  ],
  "a": [                 // ask updates
    ["29001.00", "0.8"],
    ...
  ]
}
