package tickets_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	// Intialize testing enviroment
	tickets.InitializeCache()

	/* Test cases */
	t.Run("Test invalid country name returns 0", func(t *testing.T) {
		countryName := "Wakanda"
		expectedValue := 0
		actualValue := tickets.GetTotalTickets(countryName)
		require.Equal(t, expectedValue, actualValue)
	})

	t.Run("Test function with Brazil retutns 45", func(t *testing.T) {
		countryName := "Brazil"
		expectedValue := 45
		actualValue := tickets.GetTotalTickets(countryName)
		require.Equal(t, expectedValue, actualValue)
	})
}

func TestFlyCountByTimePeriod(t *testing.T) {
	// Initialize testing enviroment
	tickets.InitializeCache()

	/* Test cases */
	t.Run("Test invalid time period returns 0", func(t *testing.T) {
		var invalidTimePeriod tickets.TimePeriod = -1
		expectedValue := 0
		actualValue, err := tickets.FlyCountByTimePeriod(invalidTimePeriod)
		require.Nil(t, err)
		require.Equal(t, expectedValue, actualValue)
	})

	t.Run("Test every fly ticket has only one time period", func(t *testing.T) {
		timePeriods := []tickets.TimePeriod{
			tickets.Madrugada,
			tickets.Mañana,
			tickets.Tarde,
			tickets.Noche,
		}

		ticketSum := 0
		for _, period := range timePeriods {
			ticketCount, err := tickets.FlyCountByTimePeriod(period)
			require.Nil(t, err)
			ticketSum += ticketCount
		}
		expectedValue := 1000
		require.Equal(t, expectedValue, ticketSum)
	})

}

func TestAverageDestination(t *testing.T) {
	// Intialize testing enviroment
	tickets.InitializeCache()

	/* Test cases */
	t.Run("Test invalid country name returns 0.0", func(t *testing.T) {
		countryName := "Wakanda"
		expectedValue := 0.0
		actualValue := tickets.AverageDestination(countryName)
		require.Equal(t, expectedValue, actualValue)
	})

	t.Run("Test function with Brazil retutns 4.5", func(t *testing.T) {
		countryName := "Brazil"
		var expectedValue float64 = 45.0 / 1000.0 * 100.0
		actualValue := tickets.AverageDestination(countryName)
		require.Equal(t, expectedValue, actualValue)
	})
}
