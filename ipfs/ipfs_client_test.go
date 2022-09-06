package ipfs

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChainClient_IPFS(t *testing.T) {
	freezer := New("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDYwRDI3M0NhYTJEQjA1MzgzZjlmNzAxMDE5OTQ1MDRFNjJiYzUyM2MiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY0Mzk4MTA0MTI4OSwibmFtZSI6Ik1hcmt5In0.POucbq990jW8wRv9L8bT68JckG1-TZ4EQc2I_VwnwXo", "https://api.nft.storage/upload")
	file, err := os.Open("./testdata/ipfs.json")
	require.Nil(t, err)

	resourceUrl, metadataUrl, err := freezer.StoreWithMeta("dd", "22", file)
	require.Nil(t, err)

	fmt.Println(resourceUrl)
	fmt.Println(metadataUrl)

}
