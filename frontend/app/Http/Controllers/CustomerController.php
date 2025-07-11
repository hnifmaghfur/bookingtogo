<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;

class CustomerController extends Controller
{
    private $apiBaseUrl = 'http://localhost:8050'; // Ganti dengan URL API Anda

    public function create()
    {
        try {
            // Ambil data nationalities dari API
            $response = Http::get($this->apiBaseUrl . '/nationalities');
            $nationalities = $response->successful() ? $response->json() : [];

            return view('customer.create', compact('nationalities'));
        } catch (\Exception $e) {
            Log::error('Error fetching nationalities: ' . $e->getMessage());
            return view('customer.create', ['nationalities' => []]);
        }
    }

    public function store(Request $request)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'email' => 'required|email|max:255',
            'phone' => 'required|string|max:20',
            'date_of_birth' => 'required|date',
            'nationality' => 'required|string',
            'family_list' => 'nullable|array',
            'family_list.*.name' => 'required_with:family_list|string|max:255',
            'family_list.*.date_of_birth' => 'required_with:family_list|date',
            'family_list.*.relation' => 'required_with:family_list|string|max:100',
        ]);

        try {
            // Simpan data customer ke API
            $customerData = [
                'cst_name' => $request->name,
                'cst_email' => $request->email,
                'cst_phonenum' => $request->phone,
                'cst_dob' => $request->date_of_birth,
                'nationality_id' => (int) $request->nationality,
            ];

            Log::info('Customer data: ' . json_encode($customerData));

            $customerResponse = Http::post($this->apiBaseUrl . '/customers', $customerData);

            if (!$customerResponse->successful()) {
                Log::error('Error saving customer: ' . $customerResponse->json());
                throw new \Exception(json_encode($customerResponse->json()));
            }

            Log::info('Customer data: ' . json_encode($customerData));

            $customerId = $customerResponse->json()['cst_id'] ?? null;

            // Simpan family list jika ada
            if ($request->has('family_list') && is_array($request->family_list)) {
                $bulkFamilyData = [];
                foreach ($request->family_list as $family) {
                    if (!empty($family['name']) && !empty($family['date_of_birth']) && !empty($family['relation'])) {
                        $bulkFamilyData[] = [
                            'cst_id' => $customerId,
                            'fl_name' => $family['name'],
                            'fl_dob' => $family['date_of_birth'],
                            'fl_relation' => $family['relation'],
                        ];
                    }
                }
                Log::info('Bulk family data: ' . json_encode($bulkFamilyData));
                if (count($bulkFamilyData) > 0) {
                    $familyResponse = Http::post($this->apiBaseUrl . '/family-lists', $bulkFamilyData);
                    if (!$familyResponse->successful()) {
                        Log::error('Error saving family (bulk): ' . json_encode($familyResponse->json()));
                        throw new \Exception(json_encode($familyResponse->json()));
                    }
                }
            }

            return redirect()->route('customer.create')->with('success', 'Customer berhasil disimpan!');
        } catch (\Exception $e) {
            Log::error('Error saving customer: ' . $e->getMessage());
            return redirect()->back()->with('error', 'Terjadi kesalahan saat menyimpan data.')->withInput();
        }
    }
}
