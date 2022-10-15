const weights = [{ host: 'host1', weight: 100 }, { host: 'host2', weight: 200 }, { host: 'host3', weight: 300 }]

const actual_weights = [0, 0, 0]



while (true) {

    totalWeights = 100 +200 +300
    actual_weights = [100, 200, 300];

    actual_weights = [100, 200, 300 - totalWeights];//1
    actual_weights = [200, 400, 0];//add weight
    actual_weights = [200, 400 - totalWeights, 0];//2
    actual_weights = [300, 0, 300];//add weight
    actual_weights = [300 - totalWeights, 200, 300];//3

    return max(weights);
}