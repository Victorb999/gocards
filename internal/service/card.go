package service

import "database/sql"

type Card struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TypeCard    string `json:"type_card"`
	Description string `json:"description"`
	ATK         int    `json:"atk"`
	DEF         int    `json:"def"`
}

type CardService struct {
	db *sql.DB
}

func NewCardService(db *sql.DB) *CardService {
	return &CardService{db: db}
}

func (s *CardService) CreateCard(card *Card) error {
	query := `INSERT INTO cards (name, typeCard, description, atk, def) VALUES ($1, $2, $3, $4, $5)`
	result, err := s.db.Exec(query, card.Name, card.TypeCard, card.Description, card.ATK, card.DEF)
	if err != nil {
		return err
	}
	lastInserId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	card.ID = int(lastInserId)
	return nil
}

func (s *CardService) GetCards() ([]Card, error) {
	query := `SELECT id, name, typeCard, description, atk, def FROM cards`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := []Card{}
	for rows.Next() {
		card := Card{}
		err := rows.Scan(&card.ID, &card.Name, &card.TypeCard, &card.Description, &card.ATK, &card.DEF)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (s *CardService) GetCard(id int) (*Card, error) {
	query := `SELECT id, name, typeCard, description, atk, def FROM cards WHERE id = $1`

	row := s.db.QueryRow(query, id)
	card := &Card{}

	err := row.Scan(&card.ID, &card.Name, &card.TypeCard, &card.Description, &card.ATK, &card.DEF)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *CardService) DeleteCard(id int) error {
	query := `DELETE FROM cards WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CardService) UpdateCard(id int, card *Card) error {
	query := `UPDATE cards SET name = $1, typeCard = $2, description = $3, atk = $4, def = $5 WHERE id = $6`
	_, err := s.db.Exec(query, card.Name, card.TypeCard, card.Description, card.ATK, card.DEF, id)
	if err != nil {
		return err
	}
	return nil
}
