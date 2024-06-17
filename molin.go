	query := datastore.NewQuery("Game").Filter("Active =", true).Limit(1)
	var activeGame Game
	if err := client.Get(ctx, query.KeysOnly(), &activeGame); err != nil {
		return err
	}
	return transactionBlock.Object(activeGame.Address)  
